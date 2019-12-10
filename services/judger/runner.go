package judger

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"Rabbit-OJ-Backend/utils/path"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Runner(
	sid, codePath string,
	compileInfo *config.CompileInfo,
	caseCount, timeLimit, spaceLimit, casePath, outputPath string,
) error {
	fmt.Printf("(%s) [Runner] Compile OK, start run container %s \n", sid, codePath)

	err := path.TouchFile(codePath + ".result")
	if err != nil {
		fmt.Printf("(%s) %+v \n", sid, err)
		return err
	}
	fmt.Printf("(%s) [Runner] Touched empty result file for build \n", sid)

	containerConfig := &container.Config{
		Image:           compileInfo.RunImage,
		NetworkDisabled: true,
		Env: []string{
			"EXEC_COMMAND=" + compileInfo.RunArgs,
			"CASE_COUNT=" + caseCount,
			"TIME_LIMIT=" + timeLimit,
			"SPACE_LIMIT=" + spaceLimit,
			"Role=Tester",
		},
	}

	containerHostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			//	{
			//		Source:   codePath + ".o",
			//		Target:   compileInfo.BuildTarget,
			//		ReadOnly: true,
			//		Type:     mount.TypeBind,
			//	},
			//	{
			//		Source:   "/app/server",
			//		Target:   "/app/tester",
			//		ReadOnly: true,
			//		Type:     mount.TypeBind,
			//	},
			//	{
			//		Source: codePath + ".result",
			//		Target: "/result/info.json",
			//		Type:   mount.TypeBind,
			//	},
			{
				Source:   casePath,
				Target:   "/case",
				ReadOnly: true,
				Type:     mount.TypeBind,
			},
			//	{
			//		Source: outputPath,
			//		Target: "/output",
			//		Type:   mount.TypeBind,
			//	},
		},
		Binds: []string{
			path.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
			path.DockerHostConfigBinds(codePath+".result", "/result/info.json"),
			path.DockerHostConfigBinds(outputPath, "/output"),
		},
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

	fmt.Printf("(%s) [Runner] Running container \n", sid)
	if err := docker.Client.ContainerStart(docker.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("(%s) [Runner] %+v \n", sid, err)
		return err
	}

	statusCh, errCh := docker.Client.ContainerWait(docker.Context, resp.ID, container.WaitConditionNotRunning)
	fmt.Printf("(%s) [Runner] Waiting for status \n", sid)
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Printf("(%s) %+v \n", sid, status)
	case <-time.After(120 * time.Second):
		return errors.New("run timeout")
	}

	return nil
}
