package judger

import "fmt"

func StartMachine(index uint, queueChan chan []byte) {
	fmt.Printf("[Machine] Concurrent #%d started \n", index)
	for {
		body := <-queueChan

		fmt.Printf("[Machine] #%d machine START \n", index)
		okChan := make(chan bool)
		go JudgeRequestBridge(body, okChan)
		select {
		case <-okChan:
			close(okChan)
		}
		fmt.Printf("[Machine] #%d machine FINISH \n", index)
	}
}
