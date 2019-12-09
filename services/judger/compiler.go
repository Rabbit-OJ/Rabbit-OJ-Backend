package judger

import (
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
		AutoRemove: true,
		Binds: []string{
			utils.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
		},
	}

	fmt.Printf("(%s) [Compile] Creating container \n", sid)
	resp, err := DockerClient.ContainerCreate(DockerContext,
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

	fmt.Printf("(%s) [Compile] Running container \n", sid)
	if err := DockerClient.ContainerStart(DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}

	statusCh, errCh := DockerClient.ContainerWait(DockerContext, resp.ID, container.WaitConditionNotRunning)
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
