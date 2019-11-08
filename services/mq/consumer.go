package mq

import (
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go JudgeStart(&delivery)
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go JudgeResultStart(&delivery)
	}
}