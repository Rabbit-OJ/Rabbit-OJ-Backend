package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	queueChan := make(chan []byte)

	for i := uint(0); i < config.Global.Concurrent.JudgeCount; i++ {
		go StartMachine(i, queueChan)
	}

	for delivery := range deliveries {
		queueChan <- delivery.Body
		// block until one machine receive the request body, then ACK
		_ = delivery.Ack(false)
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
