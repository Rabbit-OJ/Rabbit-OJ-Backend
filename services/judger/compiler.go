package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Compiler(codePath, language string) error {
	compileOptions, ok := utils.CompileObject[language]
	if !ok {
		return errors.New("language doesn't exist")
	}

	containerConfig := &container.Config{
		Entrypoint:      []string{compileOptions.Args},
		Image:           compileOptions.Image,
		NetworkDisabled: true,
		StopTimeout:     &compileOptions.CompileTime,
	}
	containerHostConfig := &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Source:   codePath,
				Target:   "/submit/code.cpp",
				ReadOnly: true,
			},
			{
				Source: codePath + ".o",
				Target: "/compile/code.o",
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
		if err != nil {
			return err
		}
	case status := <-statusCh:
		fmt.Println(status)
	case <-time.After(time.Duration(compileOptions.CompileTime) * time.Second):
		return errors.New("compile timeout")
	}

	return nil
}
