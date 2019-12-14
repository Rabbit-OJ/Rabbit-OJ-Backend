package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	queueChan := make(chan *amqp.Delivery)

	for i := uint(0); i < config.Global.Concurrent.Judge; i++ {
		go StartMachine(i, queueChan)
	}

	for delivery := range deliveries {
		queueChan <- &delivery
		// block until one machine receive the request body, then ACK
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		go JudgeResponseBridge(&delivery)
	}
}
