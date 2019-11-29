package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Compiler(codePath string, compileInfo *utils.CompileInfo) error {
	fmt.Println("[Compile] Start" + codePath)

	err := utils.TouchFile(codePath + ".o")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("[Compile] Touched empty output file for build")
	containerConfig := &container.Config{
		Entrypoint:      []string{compileInfo.BuildArgs},
		Image:           compileInfo.BuildImage,
		NetworkDisabled: true,
		StopTimeout:     &compileInfo.BuildTime,
	}

	containerHostConfig := &container.HostConfig{
		//AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Source:   codePath,
				Target:   compileInfo.BuildSource,
				ReadOnly: true,
				Type:     mount.TypeBind,
			},
			{
				Source: codePath + ".o",
				Target: compileInfo.BuildTarget,
				Type:     mount.TypeBind,
			},
		},
	}

	fmt.Println("[Compile] Creating container")
	resp, err := DockerClient.ContainerCreate(DockerContext,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	fmt.Println("[Compile] Running container")
	if err := DockerClient.ContainerStart(DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}

	statusCh, errCh := DockerClient.ContainerWait(DockerContext, resp.ID, container.WaitConditionNotRunning)
	fmt.Println("[Compile] Waiting for status")
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Println(status)
	case <-time.After(time.Duration(compileInfo.BuildTime) * time.Second):
		return errors.New("compile timeout")
	}

	return nil
}
