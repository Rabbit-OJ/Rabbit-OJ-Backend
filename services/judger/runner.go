package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Runner(
	codePath string,
	compileInfo *utils.CompileInfo,
	caseCount, timeLimit, spaceLimit, casePath, outputPath string,
) error {
	fmt.Println("[Runner] Compile OK, start run container " + codePath)

	err := utils.TouchFile(codePath + ".result")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("[Runner] Touched empty result file for build")

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
		AutoRemove: true,
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
			utils.DockerHostConfigBinds(codePath+".o", compileInfo.BuildTarget),
			utils.DockerHostConfigBinds(codePath+".result", "/result/info.json"),
			utils.DockerHostConfigBinds(outputPath, "/output"),
		},
	}

	fmt.Println("[Runner] Creating container")
	resp, err := DockerClient.ContainerCreate(DockerContext,
		containerConfig,
		containerHostConfig,
		nil,
		"")

	if err != nil {
		return err
	}

	fmt.Println("[Runner] Running container")
	if err := DockerClient.ContainerStart(DockerContext, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}

	statusCh, errCh := DockerClient.ContainerWait(DockerContext, resp.ID, container.WaitConditionNotRunning)
	fmt.Println("[Runner] Waiting for status")
	select {
	case err := <-errCh:
		return err
	case status := <-statusCh:
		fmt.Println(status)
	case <-time.After(120 * time.Second):
		return errors.New("run timeout")
	}

	return nil
}
