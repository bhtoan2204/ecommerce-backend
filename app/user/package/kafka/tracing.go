package kafka

import (
	"context"
	"fmt"
	"strconv"
	"user/package/tracer"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func (p *producer) startSpan(ctx context.Context, msg *kafka.Message) (context.Context, trace.Span) {
	carrier := NewMessageCarrier(msg)
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	namespan := fmt.Sprintf("%s send", *msg.TopicPartition.Topic)
	opts := p.buildSpanOpts(msg)

	return tracer.StartSpan(ctx, namespan, opts...)
}

func (c *consumer) startSpan(msg *kafka.Message) (context.Context, trace.Span) {

	carrier := NewMessageCarrier(msg)
	ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)

	namespan := fmt.Sprintf("%s receive", *msg.TopicPartition.Topic)
	opts := c.buildSpanOpts(msg)

	return tracer.StartSpan(ctx, namespan, opts...)
}

func (c *producer) buildSpanOpts(msg *kafka.Message) []trace.SpanStartOption {
	result := []trace.SpanStartOption{}
	offset := strconv.FormatInt(int64(msg.TopicPartition.Offset), 10)

	result = append(result,
		trace.WithAttributes(
			semconv.MessagingDestinationNameKey.String(*msg.TopicPartition.Topic),
			semconv.MessagingMessageIDKey.String(offset),
			semconv.MessagingKafkaMessageKeyKey.String(string(msg.Key)),
			semconv.MessagingKafkaSourcePartitionKey.Int64(int64(msg.TopicPartition.Partition)),
		),
		trace.WithSpanKind(trace.SpanKindProducer),
	)

	return result
}

func (c *consumer) buildSpanOpts(msg *kafka.Message) []trace.SpanStartOption {
	result := []trace.SpanStartOption{}
	offset := strconv.FormatInt(int64(msg.TopicPartition.Offset), 10)

	result = append(result,
		trace.WithAttributes(
			semconv.MessagingSourceNameKey.String(*msg.TopicPartition.Topic),
			semconv.MessagingMessageIDKey.String(offset),
			semconv.MessagingKafkaMessageKeyKey.String(string(msg.Key)),
			semconv.MessagingKafkaSourcePartitionKey.Int64(int64(msg.TopicPartition.Partition)),
		),
		trace.WithSpanKind(trace.SpanKindConsumer),
	)

	return result
}
