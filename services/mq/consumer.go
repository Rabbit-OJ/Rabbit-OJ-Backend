package mq

import (
	"Rabbit-OJ-Backend/services/config"
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func CreateJudgeRequestConsumer(topics []string, group string) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = Version
	consumer := JudgeRequestConsumer{
		ready: make(chan bool, 0),
	}

	client, err := sarama.NewConsumerGroup(config.Global.Kafka.Brokers, group, saramaConfig)
	if err != nil {
		log.Panicf("Error when creating consumer group: %v", err)
		return
	}

	ctx, _ := context.WithCancel(CancelCtx)
	go func() {
		for {
			fmt.Println("[MQ] topic: request consumer group running", group)

			if err := client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				log.Panicf("Error from ctx: %v", ctx.Err())
				return
			}

			consumer.ready = make(chan bool, 0)
		}
	}()
}

func CreateJudgeResponseConsumer(topics []string, group string) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = Version
	consumer := JudgeResponseConsumer{
		ready: make(chan bool, 0),
	}

	client, err := sarama.NewConsumerGroup(config.Global.Kafka.Brokers, group, saramaConfig)
	if err != nil {
		log.Panicf("Error when creating consumer group: %v", err)
		return
	}

	ctx, _ := context.WithCancel(CancelCtx)
	go func() {
		for {
			fmt.Println("[MQ] topic: response consumer group running", group)

			if err := client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				log.Panicf("Error from ctx: %v", ctx.Err())
				return
			}

			consumer.ready = make(chan bool, 0)
		}
	}()
}
