package docker

import (
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func PullImage(tag string) {
	fmt.Println("[DIND] pulling image : " + tag)
	out, err := DockerClient.ImagePull(DockerContext, tag, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer func() { _ = out.Close() }()

	if _, err := io.Copy(os.Stderr, out); err != nil {
		fmt.Println(err)
	}
}

func BuildImage(tag string) {
	fmt.Println("[DIND] building image from local Dockerfile : " + tag)

	dockerFileBytes, err := utils.ReadFileBytes(fmt.Sprintf("./dockerfiles/%s/Dockerfile", tag))
	if err != nil {
		panic(err)
	}
	serverFileBytes, err := utils.ReadFileBytes("./server")
	if err != nil {
		panic(err)
	}
	tarBytes, err := utils.ConvertToTar([]utils.TarFileBasicInfo{
		{
			Name: "Dockerfile",
			Body: dockerFileBytes,
		},
		{
			Name: "server",
			Body: serverFileBytes,
		},
	})
	if err != nil {
		panic(err)
	}

	resp, err := DockerClient.ImageBuild(DockerContext, tarBytes, types.ImageBuildOptions{
		Tags:   []string{tag},
		Remove: true,
	})
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stderr, resp.Body); err != nil {
		fmt.Println(err)
	}
}