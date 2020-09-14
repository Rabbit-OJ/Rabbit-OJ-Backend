package judger

import (
	"context"
	"fmt"
	"sync"
)

var (
	MachineWaitGroup sync.WaitGroup
)

func StartMachine(ctx context.Context, index uint, queueChan chan []byte) {
	fmt.Printf("[Machine] Concurrent #%d started \n", index)
	MachineWaitGroup.Add(1)
	defer MachineWaitGroup.Done()

	for {
		select {
		case delivery := <-queueChan:
			fmt.Printf("[Machine] #%d machine START \n", index)
			okChan := make(chan bool)
			go JudgeRequestBridge(delivery, okChan)
			select {
			case <-okChan:
				close(okChan)
			}
			fmt.Printf("[Machine] #%d machine FINISH \n", index)
		case <-ctx.Done():
			fmt.Printf("[Machine] #%d machine Exited \n", index)
			return
		}
	}
}
