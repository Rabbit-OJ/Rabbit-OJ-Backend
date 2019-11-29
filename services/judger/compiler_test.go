package judger

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"testing"
)

func TestCompiler(t *testing.T) {
	DockerInit()

	//containerConfig := &container.Config{
	//	Entrypoint:      []string{"sh"},
	//	Image:           "alpine:latest",
	//	OpenStdin:       true,
	//	Tty:             true,
	//	NetworkDisabled: true,
	//}
	//
	//containerHostConfig := &container.HostConfig{}

	io, err := os.OpenFile("/Users/yangziyue/1.tar", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	//resp, err := DockerClient.ContainerCreate(DockerContext,
	//	containerConfig,
	//	containerHostConfig,
	//	nil,
	//	"")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	t.Fail()
	//	return
	//}
	//
	//if err := DockerClient.ContainerStart(DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	fmt.Println(err)
	//	t.Fail()
	//	return
	//}

	if err := DockerClient.CopyToContainer(DockerContext, "b542af23d55c", "/home", io, types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
		CopyUIDGID:                false,
	}); err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
}
