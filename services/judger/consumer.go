package judger

import (
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go JudgeRequestBridge(&delivery)
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		_ = delivery.Ack(false)

		go JudgeResponseBridge(&delivery)
	}
}