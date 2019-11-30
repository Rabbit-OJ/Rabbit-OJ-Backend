package judger

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"testing"
)

func TestCompiler(t *testing.T) {
	DockerInit()

	containerConfig := &container.Config{
		WorkingDir:      "/home",
		Cmd:             []string{"gcc", "--version"},
		Image:           "gcc:9.2.0",
		NetworkDisabled: true,
	}

	containerHostConfig := &container.HostConfig{}

	resp, err := DockerClient.ContainerCreate(DockerContext,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if err := DockerClient.ContainerStart(DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

}
