package mq

import (
	"github.com/streadway/amqp"
)

func TestCaseHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go TestCaseConsumer(&delivery)
	}
}

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go judgeStart(&delivery)
	}
}
