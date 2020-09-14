package mq

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func RequeueHandler() {
	for buf := range JudgeRequeueDeliveryChan {
		msg := &sarama.ProducerMessage{
			Topic: config.JudgeRequestTopicName,
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", time.Now().UnixNano())),
			Value: sarama.StringEncoder(buf),
		}

		AsyncProducer.Input() <- msg
	}
}
