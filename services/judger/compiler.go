package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"time"
)

func Compiler(codePath string, code []byte, compileInfo *utils.CompileInfo) error {
	fmt.Println("[Compile] Start" + codePath)

	err := utils.TouchFile(codePath + ".o")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("[Compile] Touched empty output file for build")
	containerConfig := &container.Config{
		Entrypoint:      compileInfo.BuildArgs,
		//Entrypoint:      []string{"bash"},
		Tty:             true,
		OpenStdin:       true,
		Image:           compileInfo.BuildImage,
		NetworkDisabled: true,
		StopTimeout:     &compileInfo.BuildTime,
	}

	containerHostConfig := &container.HostConfig{
		AutoRemove: true,
		Binds: []string{
			utils.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
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

	fmt.Println("[Compile] Copying files to container")
	io, err := utils.ConvertToTar([]utils.TarFileBasicInfo{{compileInfo.SourceFileName, code}})
	if err != nil {
		return err
	}

	if err := DockerClient.CopyToContainer(
		DockerContext,
		resp.ID,
		compileInfo.ExecFilePath,
		io,
		types.CopyToContainerOptions{
			AllowOverwriteDirWithFile: true,
			CopyUIDGID:                false,
		}); err != nil {
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
