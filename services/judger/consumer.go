package judger

import (
	"github.com/streadway/amqp"
)

// We should NOT judge multiple codes at the same time, or the time will not accurate

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		okChan := make(chan bool)

		_ = delivery.Ack(false)
		go JudgeRequestBridge(&delivery, okChan)

		select {
		case <-okChan:
			close(okChan)
		}
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		okChan := make(chan bool)

		_ = delivery.Ack(false)
		go JudgeResponseBridge(&delivery, okChan)

		select {
		case <-okChan:
			close(okChan)
		}
	}
}
