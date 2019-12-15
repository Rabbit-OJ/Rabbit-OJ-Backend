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

	for {
		select {
		case delivery := <-deliveries:
			queueChan <- &delivery
		case <-MachineContext.Done():
			return
		}
	}
}

func JudgeResultHandler(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		go JudgeResponseBridge(&delivery)
	}
}
