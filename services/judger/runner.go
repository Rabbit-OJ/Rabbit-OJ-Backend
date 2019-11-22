package judger

import (
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"time"
)

func Runner(
	codePath string,
	compileInfo *utils.CompileInfo,
	caseCount, timeLimit, spaceLimit, casePath, outputPath string,
) error {
	fmt.Println("Compile OK, start run container " + codePath)

	containerConfig := &container.Config{
		Entrypoint:      []string{"/app/tester"},
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
			{
				Source:   codePath + ".o",
				Target:   compileInfo.BuildTarget,
				ReadOnly: true,
			},
			{
				Source:   "/app/server",
				Target:   "/app/tester",
				ReadOnly: true,
			},
			{
				Source: codePath + ".result",
				Target: "/result/info.json",
			},
			{
				Source:   casePath,
				Target:   "/case",
				ReadOnly: true,
			},
			{
				Source: outputPath,
				Target: "/output",
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
		return err
	case status := <-statusCh:
		fmt.Println(status)
	case <-time.After(120 * time.Second):
		return errors.New("run timeout")
	}

	return nil
}
