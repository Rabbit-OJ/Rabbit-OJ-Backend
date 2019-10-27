package mq

import (
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go judgeStart(&delivery)
	}
}
