package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"github.com/streadway/amqp"
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	queueChan := make(chan []byte)

	for i := uint(0); i < config.Global.Concurrent.Judge; i++ {
		go StartMachine(i, queueChan)
	}

	for delivery := range deliveries {
		queueChan <- delivery.Body
		// block until one machine receive the request body, then ACK
		if err := delivery.Ack(false); err != nil {
			fmt.Println(err)
		}
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		if err := delivery.Ack(false); err != nil {
			fmt.Println(err)
		}

		go JudgeResponseBridge(&delivery)
	}
}
