package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	path2 "Rabbit-OJ-Backend/utils/path"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"path"
	"time"
)

func Compiler(sid, codePath string, code []byte, compileInfo *config.CompileInfo) error {
	fmt.Printf("(%s) [Compile] Start %s \n", sid, codePath)

	err := path2.TouchFile(codePath + ".o")
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
	}

	containerHostConfig := &container.HostConfig{
		Binds: []string{
			path2.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
		},
	}

	if config.Global.AutoRemove.Containers {
		containerHostConfig.AutoRemove = true
	}

	fmt.Printf("(%s) [Compile] Creating container \n", sid)
	resp, err := docker.Client.ContainerCreate(docker.Context,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	fmt.Printf("(%s) [Compile] Copying files to container \n", sid)
	io, err := path2.ConvertToTar([]path2.TarFileBasicInfo{{path.Base(compileInfo.Source), code}})
	if err != nil {
		return err
	}

	if err := docker.Client.CopyToContainer(
		docker.Context,
		resp.ID,
		path.Dir(compileInfo.Source),
		io,
		types.CopyToContainerOptions{
			AllowOverwriteDirWithFile: true,
			CopyUIDGID:                false,
		}); err != nil {
		return err
	}

	fmt.Printf("(%s) [Compile] Running container \n", sid)
	if err := docker.Client.ContainerStart(docker.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}

	statusCh, errCh := docker.Client.ContainerWait(docker.Context, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%s) [Compile] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Printf("(%s) %+v \n", sid, status)
	case <-time.After(time.Duration(compileInfo.BuildTimeout) * time.Second):
		go docker.ForceContainerRemove(resp.ID)
		return errors.New("compile timeout")
	}

	return nil
}
