package judger

import (
	"fmt"
	"github.com/streadway/amqp"
)

func StartMachine(index uint, queueChan chan *amqp.Delivery) {
	fmt.Printf("[Machine] Concurrent #%d started \n", index)
	for {
		delivery := <-queueChan

		fmt.Printf("[Machine] #%d machine START \n", index)
		okChan := make(chan bool)
		go JudgeRequestBridge(delivery, okChan)
		select {
		case <-okChan:
			close(okChan)
		}
		fmt.Printf("[Machine] #%d machine FINISH \n", index)
	}
}
