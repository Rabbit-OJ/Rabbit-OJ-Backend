package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
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
