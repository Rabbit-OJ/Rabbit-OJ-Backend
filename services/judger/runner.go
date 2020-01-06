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
	sid uint32, codePath string,
	compileInfo *config.CompileInfo,
	caseCount, timeLimit, spaceLimit, casePath, outputPath string,
	code []byte,
) error {
	vmPath := codePath + "vm/"
	fmt.Printf("(%d) [Runner] Compile OK, start run container \n", sid)

	err := files.TouchFile(codePath + "result.json")
	if err != nil {
		fmt.Printf("(%d) %+v \n", sid, err)
		return err
	}
	fmt.Printf("(%d) [Runner] Touched empty result file for build \n", sid)

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
			files.DockerHostConfigBinds(codePath+"result.json", "/result/info.json"),
			files.DockerHostConfigBinds(outputPath, "/output"),
		},
	}
	// todo : limit memory usage

	if !compileInfo.NoBuild {
		containerHostConfig.Binds = append(containerHostConfig.Binds,
			files.DockerHostConfigBinds(vmPath, path.Dir(compileInfo.BuildTarget)))
	}

	if config.Global.AutoRemove.Containers {
		containerHostConfig.AutoRemove = true
	}

	fmt.Printf("(%d) [Runner] Creating container \n", sid)
	resp, err := docker.Client.ContainerCreate(docker.Context,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	if compileInfo.NoBuild {
		fmt.Printf("(%d) [Runner] Copying files to container \n", sid)
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

	fmt.Printf("(%d) [Runner] Running container \n", sid)
	if err := docker.Client.ContainerStart(docker.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%d) [Runner] %+v \n", sid, err)
		return err
	}

	docker.ContainerErrToStdErr(resp.ID)
	statusCh, errCh := docker.Client.ContainerWait(docker.Context, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%d) [Runner] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Printf("(%d) %+v \n", sid, status)
	case <-time.After(time.Duration(compileInfo.Constraints.RunTimeout) * time.Second):
		docker.ForceContainerRemove(resp.ID)
		return errors.New("run timeout")
	}

	return nil
}
