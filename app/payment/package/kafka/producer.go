package kafka

import (
	"context"
	"encoding/json"
	"payment/package/logger"
	"strconv"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

var _ Producer = (*producer)(nil)

type Producer interface {
	Produce(ctx context.Context, topic string, v interface{}) error
	ProduceWithKey(ctx context.Context, topic string, key []byte, v interface{}) error
	ProduceRaw(ctx context.Context, topic string, data []byte) error
	ProduceRawWithKey(ctx context.Context, topic string, key []byte, data []byte) error
	SyncProduceWithKey(ctx context.Context, topic string, key []byte, v interface{}) error

	Close(ctx context.Context)
}

type producer struct {
	ins              *kafka.Producer
	chanStop         chan bool
	startClosingOnce sync.Once
}

func NewProducer(cfg *Config) (Producer, error) {
	// https://github.com/confluentinc/librdkafka/tree/master/CONFIGURATION.md
	cfgMap := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Servers,
	}
	p, err := kafka.NewProducer(cfgMap)
	if err != nil {
		return nil, err
	}

	producer := &producer{
		ins:      p,
		chanStop: make(chan bool, 1),
	}

	go producer.listenDefaultEvent()

	return producer, nil
}

func (p *producer) listenDefaultEvent() {
	log := logger.DefaultLogger().Named("KafkaProducer")

loop:
	for {
		select {
		case <-p.chanStop:
			log.Info("Stop listen the default events channel - kafka producer")
			break loop
		case e := <-p.ins.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				// The message delivery report, indicating success or
				// permanent failure after retries have been exhausted.
				// Application level retries won't help since the client
				// is already configured to do that.
				m := ev
				if m.TopicPartition.Error != nil {
					log.Warn("Delivery failed: %v", zap.Error(m.TopicPartition.Error))
				} else {
					log.Info("Delivered message",
						zap.String("topic", *m.TopicPartition.Topic),
						zap.Int32("partition", m.TopicPartition.Partition),
						zap.Any("offset", m.TopicPartition.Offset))
				}
			default:
			}
		}
	}
}

func (p *producer) ProduceRaw(ctx context.Context, topic string, data []byte) error {
	key := []byte(strconv.Itoa(int(time.Now().Unix())))
	return p.ProduceRawWithKey(ctx, topic, key, data)
}

func (p *producer) Produce(ctx context.Context, topic string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return p.ProduceRaw(ctx, topic, data)
}

func (p *producer) ProduceWithKey(ctx context.Context, topic string, key []byte, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return p.ProduceRawWithKey(ctx, topic, key, data)
}

func (p *producer) ProduceRawWithKey(ctx context.Context, topic string, key []byte, data []byte) error {
	log := logger.FromContext(ctx)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: data,
	}

	_, span := p.startSpan(ctx, msg)
	defer span.End()

	err := p.ins.Produce(msg, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Error("producer ProduceRawWithKey got error",
			zap.Error(err), zap.ByteString("key", key), zap.ByteString("body", data))
		return err
	}

	return nil
}

func (p *producer) Close(ctx context.Context) {
	log := logger.DefaultLogger().Named("KafkaProducer")

	p.startClosingOnce.Do(func() {
		log.Info("Stoping kafka producer")
		for p.ins.Flush(5000) > 0 {
			log.Info("Still waiting to flush outstanding messages")
		}
		p.ins.Close()
		p.chanStop <- true
		log.Info("Kafka producer stopped")
	})
}

func (p *producer) SyncProduceWithKey(ctx context.Context, topic string, key []byte, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return p.SyncProduceRawWithKey(ctx, topic, key, data)
}

func (p *producer) SyncProduceRawWithKey(ctx context.Context, topic string, key []byte, data []byte) error {
	deliveryChan := make(chan kafka.Event)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: data,
	}

	_, span := p.startSpan(ctx, msg)
	defer span.End()

	err := p.ins.Produce(msg, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	msg = e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		span.RecordError(msg.TopicPartition.Error)
		span.SetStatus(codes.Error, msg.TopicPartition.Error.Error())
		return msg.TopicPartition.Error
	}

	return nil
}
