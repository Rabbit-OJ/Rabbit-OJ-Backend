package docker

import (
	"Rabbit-OJ-Backend/utils/files"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

func PullImage(tag string) {
	fmt.Println("[DIND] pulling image : " + tag)
	out, err := Client.ImagePull(Context, tag, types.ImagePullOptions{})
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

	dockerFileBytes, err := files.ReadFileBytes(fmt.Sprintf("./dockerfiles/%s/Dockerfile", tag))
	if err != nil {
		panic(err)
	}
	serverFileBytes, err := files.ReadFileBytes("./server")
	if err != nil {
		panic(err)
	}
	tarBytes, err := files.ConvertToTar([]files.TarFileBasicInfo{
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

	resp, err := Client.ImageBuild(Context, tarBytes, types.ImageBuildOptions{
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