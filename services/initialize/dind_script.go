package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"os"
	"os/exec"
)

func DindScript() {
	if config.Global.Extensions.Dind {
		fmt.Println("[DIND] exec docker command")
		cmd := exec.Command("sh -c /usr/local/bin/docker-entrypoint.sh")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			if os.Getenv("ENV") == "production" {
				panic(err)
			}
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println(err)
			if os.Getenv("ENV") == "production" {
				panic(err)
			}
		}
	}
}
