package judger

import (
	"fmt"
	"os/exec"
)

func DockerScript() {
	fmt.Println("[DIND] exec docker command")
	cmd := exec.Command("sh -c /usr/local/bin/docker-entrypoint.sh")

	// todo: panic when crashing
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		//panic(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		//panic(err)
	}
}
