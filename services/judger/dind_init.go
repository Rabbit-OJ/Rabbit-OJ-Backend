package judger

import (
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/docker/docker/api/types"
)

func pullImage(tag string) {
	fmt.Println("[DIND] pulling image : " + tag)
	out, err := DockerClient.ImagePull(DockerContext, tag, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer func() { _ = out.Close() }()
}

func buildImage(tag string) {
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

	if _, err := DockerClient.ImageBuild(DockerContext, tarBytes, types.ImageBuildOptions{
		Tags: []string{tag},
	}); err != nil {
		panic(err)
	}
}

func InitImages() {
	needImages := make(map[string]bool)

	for _, item := range utils.CompileObject {
		needImages[item.RunImage] = true
		needImages[item.RunImage] = true
	}

	fmt.Println("[DIND] fetching image list")
	images, err := DockerClient.ImageList(DockerContext, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("[DIND] comparing image list")
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if _, ok := needImages[tag]; ok {
				needImages[tag] = false
			}
		}
	}

	for imageTag, need := range needImages {
		if !need {
			continue
		}

		if v, ok := utils.LocalImages[imageTag]; ok && v {
			pullImage(imageTag)
		} else {
			buildImage(imageTag)
		}
	}
}
