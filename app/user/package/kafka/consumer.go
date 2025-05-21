package kafka

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
	"user/package/logger"

	"github.com/avast/retry-go"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

var _ Consumer = (*consumer)(nil)

type CallBack func(ctx context.Context, topic string, value []byte) error
type Handler func(ctx context.Context, value []byte) error

type Consumer interface {
	Read(callback CallBack)
	Stop()
	SetHandler(h Handler)
	GetHandler() Handler
	GetHandlerName() string
}

type consumer struct {
	ins *kafka.Consumer

	startOnce sync.Once

	startClosingOnce sync.Once
	chanStop         chan bool
	producer         Producer
	handler          Handler
	handlerName      string
	dlq              bool
}

func NewConsumer(cfg *Config, producer Producer) (Consumer, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// https://github.com/confluentinc/librdkafka/tree/master/CONFIGURATION.md
	cfgMap := &kafka.ConfigMap{
		"bootstrap.servers":        cfg.Servers,
		"group.id":                 cfg.Group,
		"auto.offset.reset":        cfg.OffsetReset,
		"enable.auto.offset.store": false,
	}
	c, err := kafka.NewConsumer(cfgMap)
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics(cfg.ConsumeTopic, nil)
	if err != nil {
		return nil, err
	}

	return &consumer{
		ins:         c,
		chanStop:    make(chan bool, 1),
		producer:    producer,
		handlerName: cfg.HandlerName,
		dlq:         cfg.DLQ,
	}, nil
}

func (c *consumer) Read(f CallBack) {
	c.startOnce.Do(func() {
		go func() {
			c.start(f)
		}()
	})
}

func (c *consumer) SetHandler(f Handler) {
	if c.handler == nil {
		c.handler = f
	}
}

func (c *consumer) GetHandler() Handler {
	return c.handler
}

func (c *consumer) GetHandlerName() string {
	return c.handlerName
}

func (c *consumer) start(f CallBack) {
	log := logger.DefaultLogger()

loop:
	for {
		select {
		case <-c.chanStop:
			log.Info("Caught signal stop kafa consumer, terminating ...")
			break loop
		default:
			ev := c.ins.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				var topic string
				if e.TopicPartition.Topic != nil {
					topic = *e.TopicPartition.Topic
				}

				ctx, span := c.startSpan(e)

				err := processMessageWithRetry(ctx, f, e)
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
					span.End()

					log.Error("consumer process got error",
						zap.String("topic", topic), zap.Error(err),
						zap.ByteString("val", e.Value),
						zap.ByteString("key", e.Key),
						zap.Int64("offset", int64(e.TopicPartition.Offset)))

					if c.dlq {
						c.StoreDLQ(ctx, e)
					}
				}

				_, err = c.ins.StoreMessage(e)
				if err != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, err.Error())
					log.Info("%% Error storing offset after message", zap.Any("topic partition", e.TopicPartition))
				}
				span.End()
			case kafka.Error:
				if !e.IsTimeout() {
					log.Error("Consume kafka got err", zap.Error(e))
				}
			default:
			}
		}
	}
}

func (c *consumer) Stop() {
	c.startClosingOnce.Do(func() {
		c.chanStop <- true
		time.Sleep(100 * time.Millisecond)
		c.ins.Close()
		logger.DefaultLogger().Info("Kafka consumer stopped")
	})
}

const DLQSuffix = "dlq"

func (c *consumer) StoreDLQ(ctx context.Context, msg *kafka.Message) {
	topic := *msg.TopicPartition.Topic
	c.producer.ProduceRawWithKey(ctx, GetDLQTopic(topic), msg.Key, msg.Value)
}

func GetDLQTopic(topic string) string {
	if !strings.HasSuffix(topic, DLQSuffix) {
		topic = fmt.Sprintf("%s.%s", topic, DLQSuffix)
	}

	return topic
}

func processMessageWithRetry(ctx context.Context, f CallBack, msg *kafka.Message) error {
	retryTimes := uint(3)
	topic := *msg.TopicPartition.Topic
	if strings.HasSuffix(topic, DLQSuffix) {
		retryTimes = 0
	}

	options := []retry.Option{
		retry.Attempts(retryTimes),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
		retry.MaxDelay(time.Second * 5),
	}

	return retry.Do(func() error {
		return f(ctx, topic, msg.Value)
	}, options...)
}
