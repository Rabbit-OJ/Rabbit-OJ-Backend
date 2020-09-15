package judger

import (
	"Rabbit-OJ-Backend/services/channel"
	"Rabbit-OJ-Backend/services/config"
	"context"
)

var (
	MachineContext           context.Context
	MachineContextCancelFunc context.CancelFunc
)

func JudgeRequestHandler() {
	queueChan := make(chan []byte)

	MachineContext, MachineContextCancelFunc = context.WithCancel(context.Background())
	for i := uint(0); i < config.Global.Concurrent.Judge; i++ {
		go StartMachine(MachineContext, i, queueChan)
	}

	for {
		select {
		case delivery := <-channel.JudgeRequestDeliveryChan:
			queueChan <- delivery
		case <-MachineContext.Done():
			return
		}
	}
}

