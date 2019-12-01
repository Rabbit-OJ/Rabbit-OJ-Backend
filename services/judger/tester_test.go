package judger

import (
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestTester(t *testing.T) {
	peakMemory := int64(0)

	cmd := exec.Command("/Users/yangziyue/Downloads/test2.o")
	memoryMonitorCloseChan := make(chan bool)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	errChan, successChan, memoryMonitorChan, memoryMonitorCloseChan := make(chan error), make(chan bool), make(chan bool), make(chan bool)
	defer func() {
		close(errChan)
		close(successChan)
		close(memoryMonitorChan)
		close(memoryMonitorCloseChan)
	}()

	go func(pid int) {
		for {
			select {
			case <-memoryMonitorCloseChan:
				return
			default:
				stat, err := utils.GetStat(pid)
				if err == nil {
					peakMemory = max(peakMemory,
						int64(stat.Memory/1024/1024),
					)

					fmt.Println(stat.Memory/1024/1024)
				} else {
					fmt.Println(err)
				}
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(cmd.Process.Pid)

	go func() {
		err := cmd.Wait()
		if err != nil {
			errChan <- err
		} else {
			successChan <- true
		}
	}()

	select {
	case <-successChan:
		fmt.Println("ok")
	case <-errChan:
	}

	fmt.Println(peakMemory)
}
