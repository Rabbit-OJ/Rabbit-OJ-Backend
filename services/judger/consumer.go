package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"context"
	"github.com/streadway/amqp"
)

var (
	MachineContext           context.Context
	MachineContextCancelFunc context.CancelFunc
)

func JudgeHandler(deliveries <-chan amqp.Delivery) {
	queueChan := make(chan *amqp.Delivery)

	MachineContext, MachineContextCancelFunc = context.WithCancel(context.Background())
	for i := uint(0); i < config.Global.Concurrent.Judge; i++ {
		go StartMachine(MachineContext, i, queueChan)
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
