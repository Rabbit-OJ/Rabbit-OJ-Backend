package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"context"
)

var (
	MachineContext           context.Context
	MachineContextCancelFunc context.CancelFunc

	JudgeRequestDeliveryChan  chan []byte
	JudgeResponseDeliveryChan chan []byte
)

func JudgeHandler() {
	queueChan := make(chan []byte)

	MachineContext, MachineContextCancelFunc = context.WithCancel(context.Background())
	for i := uint(0); i < config.Global.Concurrent.Judge; i++ {
		go StartMachine(MachineContext, i, queueChan)
	}

	for {
		select {
		case delivery := <-JudgeRequestDeliveryChan:
			queueChan <- delivery
		case <-MachineContext.Done():
			return
		}
	}
}

func JudgeResultHandler() {
	for delivery := range JudgeResponseDeliveryChan {
		go JudgeResponseBridge(delivery)
	}
}
