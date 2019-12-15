package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"Rabbit-OJ-Backend/utils/files"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"path"
	"time"
)

func Runner(
	sid, codePath string,
	compileInfo *config.CompileInfo,
	caseCount, timeLimit, spaceLimit, casePath, outputPath string,
	code []byte,
) error {
	fmt.Printf("(%s) [Runner] Compile OK, start run container %s \n", sid, codePath)

	err := files.TouchFile(codePath + ".result")
	if err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}
	fmt.Printf("(%s) [Runner] Touched empty result file for build \n", sid)

	containerConfig := &container.Config{
		Image:           compileInfo.RunImage,
		NetworkDisabled: true,
		Env: []string{
			"EXEC_COMMAND=" + compileInfo.RunArgsJSON,
			"CASE_COUNT=" + caseCount,
			"TIME_LIMIT=" + timeLimit,
			"SPACE_LIMIT=" + spaceLimit,
			"Role=Tester",
		},
	}

	containerHostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Source:   casePath,
				Target:   "/case",
				ReadOnly: true,
				Type:     mount.TypeBind,
			},
		},
		Binds: []string{
			files.DockerHostConfigBinds(codePath+".result", "/result/info.json"),
			files.DockerHostConfigBinds(outputPath, "/output"),
		},
	}
	// todo : limit memory usage

	if !compileInfo.NoBuild {
		containerHostConfig.Binds = append(containerHostConfig.Binds,
			files.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget))
	}

	if config.Global.AutoRemove.Containers {
		containerHostConfig.AutoRemove = true
	}

	fmt.Printf("(%s) [Runner] Creating container \n", sid)
	resp, err := docker.Client.ContainerCreate(docker.Context,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	if compileInfo.NoBuild {
		fmt.Printf("(%s) [Runner] Copying files to container \n", sid)
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
	}

	fmt.Printf("(%s) [Runner] Running container \n", sid)
	if err := docker.Client.ContainerStart(docker.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%s) [Runner] %+v \n", sid, err)
		return err
	}

	docker.ContainerErrToStdErr(resp.ID)
	statusCh, errCh := docker.Client.ContainerWait(docker.Context, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%s) [Runner] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Printf("(%s) %+v \n", sid, status)
	case <-time.After(time.Duration(compileInfo.RunTimeout) * time.Second):
		docker.ForceContainerRemove(resp.ID)
		return errors.New("run timeout")
	}

	return nil
}
