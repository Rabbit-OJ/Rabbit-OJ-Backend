package routine

import (
	"Rabbit-OJ-Backend/services/judger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RegisterSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			fmt.Printf("[Signal] Exit signal Received: %+v \n", s)

			judger.MachineContextCancelFunc()
			judger.MachineWaitGroup.Wait()
			judger.CallbackWaitGroup.Wait()

			fmt.Printf("[Signal] Wait 30 seconds for asynchronized tasks... \n")
			<-time.After(30 * time.Second)
			os.Exit(0)
		}
	}()
}
