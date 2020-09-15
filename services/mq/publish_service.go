package mq

import (
	"Rabbit-OJ-Backend/services/channel"
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"github.com/Shopify/sarama"
)

func PublishService() {
	for message := range channel.MQPublishMessageChannel {
		go func(msg *channel.MQMessage) {
			mqMessage := &sarama.ProducerMessage{
				Topic: config.JudgeRequestTopicName,
				Key:   sarama.StringEncoder(msg.Key),
				Value: sarama.StringEncoder(msg.Value),
			}

			if msg.Async {
				AsyncProducer.Input() <- mqMessage
			} else {
				if _, _, err := SyncProducer.SendMessage(mqMessage); err != nil {
					fmt.Println("[MQ] sync send error ", err)
				}
			}
		}(message)
	}
}
