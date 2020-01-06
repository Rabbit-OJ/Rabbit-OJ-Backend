package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func ForceContainerRemove(ID string) {
	err := Client.ContainerRemove(Context, ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func ContainerErrToStdErr(ID string) {
	go func() {
		out, err := Client.ContainerLogs(Context, ID, types.ContainerLogsOptions{
			ShowStderr: true,
			ShowStdout: true,
			Follow:     true,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() { _ = out.Close() }()

		if _, err := io.Copy(os.Stderr, out); err != nil {
			fmt.Println(err)
		}
	}()
}
