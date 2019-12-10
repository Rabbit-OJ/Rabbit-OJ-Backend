package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"time"
)

func Compiler(sid, codePath string, code []byte, compileInfo *utils.CompileInfo) error {
	fmt.Printf("(%s) [Compile] Start %s \n", sid, codePath)

	err := utils.TouchFile(codePath + ".o")
	if err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}

	fmt.Printf("(%s) [Compile] Touched empty output file for build \n", sid)
	containerConfig := &container.Config{
		Entrypoint:      compileInfo.BuildArgs,
		Tty:             true,
		OpenStdin:       true,
		Image:           compileInfo.BuildImage,
		NetworkDisabled: true,
		StopTimeout:     &compileInfo.BuildTime,
	}

	containerHostConfig := &container.HostConfig{
		Binds: []string{
			utils.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
		},
	}

	if config.Global.AutoRemove.Containers {
		containerHostConfig.AutoRemove = true
	}

	fmt.Printf("(%s) [Compile] Creating container \n", sid)
	resp, err := docker.DockerClient.ContainerCreate(docker.DockerContext,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	fmt.Printf("(%s) [Compile] Copying files to container \n", sid)
	io, err := utils.ConvertToTar([]utils.TarFileBasicInfo{{compileInfo.SourceFileName, code}})
	if err != nil {
		return err
	}

	if err := docker.DockerClient.CopyToContainer(
		docker.DockerContext,
		resp.ID,
		compileInfo.ExecFilePath,
		io,
		types.CopyToContainerOptions{
			AllowOverwriteDirWithFile: true,
			CopyUIDGID:                false,
		}); err != nil {
		return err
	}

	fmt.Printf("(%s) [Compile] Running container \n", sid)
	if err := docker.DockerClient.ContainerStart(docker.DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}

	statusCh, errCh := docker.DockerClient.ContainerWait(docker.DockerContext, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%s) [Compile] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Printf("(%s) %+v \n", sid, status)
	case <-time.After(time.Duration(compileInfo.BuildTime) * time.Second):
		return errors.New("compile timeout")
	}

	return nil
}
