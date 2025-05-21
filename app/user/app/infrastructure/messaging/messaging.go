package messaging

import (
	"fmt"
	"user/package/settings"

	"user/package/kafka"
)

func buildConsumers(cfg *settings.Config, mapHandler map[string]kafka.Handler, producer kafka.Producer) []kafka.Consumer {
	var consumers []kafka.Consumer

	for _, group := range cfg.Kafka.ConsumerGroups {
		if group.PrefixGroup == "" {
			continue
		}

		for _, item := range group.ConsumeTopics {
			if item.Topic == "" || item.Handler == "" {
				continue
			}

			handler, ok := mapHandler[item.Handler]
			if !ok {
				panic(fmt.Sprintf("not found handler %s", item.Handler))
			}

			for i := 0; i < item.NumberConsumer; i++ {
				c, err := kafka.NewConsumer(
					&kafka.Config{
						Servers:      cfg.Kafka.Servers,
						Group:        group.PrefixGroup,
						OffsetReset:  group.OffsetReset,
						ConsumeTopic: []string{item.Topic},
						HandlerName:  item.Handler,
						DLQ:          item.DLQ,
					},
					producer,
				)
				if err == nil {
					c.SetHandler(handler)
					consumers = append(consumers, c)
				}
			}

			for i := 0; i < item.NumberDLQConsumer; i++ {
				c, err := kafka.NewConsumer(
					&kafka.Config{
						Servers:      cfg.Kafka.Servers,
						Group:        group.PrefixGroup,
						OffsetReset:  group.OffsetReset,
						ConsumeTopic: []string{kafka.GetDLQTopic(item.Topic)},
						HandlerName:  item.Handler,
						DLQ:          item.DLQ,
					},
					producer,
				)
				if err == nil {
					c.SetHandler(handler)
					consumers = append(consumers, c)
				}
			}
		}
	}

	return consumers
}
