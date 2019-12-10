package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"fmt"
	"github.com/docker/docker/api/types"
)

func DockerImages() {
	needImages := make(map[string]bool)

	for _, item := range config.CompileObject {
		needImages[item.RunImage] = true
		needImages[item.RunImage] = true
	}

	fmt.Println("[DIND] fetching image list")
	images, err := docker.Client.ImageList(docker.Context, types.ImageListOptions{})
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

		if v, ok := config.LocalImages[imageTag]; ok && v {
			docker.PullImage(imageTag)
		} else {
			docker.BuildImage(imageTag)
		}
	}
}
