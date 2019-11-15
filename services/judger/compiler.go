package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Compiler(codePath string, compileInfo *utils.CompileInfo) error {
	containerConfig := &container.Config{
		Entrypoint:      []string{compileInfo.BuildArgs},
		Image:           compileInfo.BuildImage,
		NetworkDisabled: true,
		StopTimeout:     &compileInfo.BuildTime,
	}

	containerHostConfig := &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Source:   codePath,
				Target:   compileInfo.BuildSource,
				ReadOnly: true,
			},
			{
				Source: codePath + ".o",
				Target: compileInfo.BuildTarget,
			},
		},
	}

	resp, err := DockerClient.ContainerCreate(DockerContext,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	statusCh, errCh := DockerClient.ContainerWait(DockerContext, resp.ID, container.WaitConditionNotRunning)
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
