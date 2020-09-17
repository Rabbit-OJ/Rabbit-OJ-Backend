package routine

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/Rabbit-OJ/Rabbit-OJ-Judger"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for _ = range c {
			judger.MachineContextCancelFunc()
			config.CancelGlobalContext()
			judger.MachineWaitGroup.Wait()
			judger.CallbackWaitGroup.Wait()
			os.Exit(0)
		}
	}()
}
