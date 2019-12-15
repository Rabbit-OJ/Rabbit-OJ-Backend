package routine

import (
	"Rabbit-OJ-Backend/services/judger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			fmt.Printf("[Signal] Received: %+v \n", s)

			judger.MachineContextCancelFunc()
			judger.MachineWaitGroup.Wait()
			judger.CallbackWaitGroup.Wait()

			os.Exit(0)
		}
	}()
}
