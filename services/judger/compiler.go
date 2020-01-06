package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"Rabbit-OJ-Backend/utils/files"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"os"
	"path"
	"time"
)

func Compiler(sid uint32, codePath string, code []byte, compileInfo *config.CompileInfo) error {
	vmPath := codePath + "vm/"
	fmt.Printf("(%d) [Compile] Start %s \n", sid, codePath)

	fmt.Printf("(%d) [Compile] Touched empty output file for build \n", sid)
	containerConfig := &container.Config{
		Entrypoint:      compileInfo.BuildArgs,
		Tty:             true,
		OpenStdin:       true,
		Image:           compileInfo.BuildImage,
		NetworkDisabled: true,
	}

	containerHostConfig := &container.HostConfig{
		Binds: []string{
			files.DockerHostConfigBinds(vmPath, path.Dir(compileInfo.BuildTarget)),
		},
		Resources: container.Resources{
			NanoCPUs: compileInfo.Constraints.CPU,
			Memory:   compileInfo.Constraints.Memory,
		},
	}

	if config.Global.AutoRemove.Containers {
		containerHostConfig.AutoRemove = true
	}

	fmt.Printf("(%d) [Compile] Creating container \n", sid)
	resp, err := docker.Client.ContainerCreate(docker.Context,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	fmt.Printf("(%d) [Compile] Copying files to container \n", sid)
	io, err := files.ConvertToTar([]files.TarFileBasicInfo{{path.Base(compileInfo.Source), code}})
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

	fmt.Printf("(%d) [Compile] Running container \n", sid)
	if err := docker.Client.ContainerStart(docker.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%d) %+v \n", sid, err)
		return err
	}

	docker.ContainerErrToStdErr(resp.ID)
	statusCh, errCh := docker.Client.ContainerWait(docker.Context, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%d) [Compile] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		if err := checkBuildResult(vmPath + path.Base(compileInfo.BuildTarget)); err != nil {
			return err
		}
		fmt.Printf("(%d) %+v \n", sid, status)
		break
	case <-time.After(time.Duration(compileInfo.Constraints.BuildTimeout) * time.Second):
		go docker.ForceContainerRemove(resp.ID)
		return errors.New("compile timeout")
	}

	return nil
}

func checkBuildResult(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}

	if file.Size() <= int64(len(files.MagicBytes)) {
		return errors.New("compile file invalid")
	}

	return nil
}
