package docker

import (
	"Rabbit-OJ-Backend/services/config"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func InitDocker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	Context, Client = ctx, cli
	if config.Global.Extensions.AutoPull {
		InitDockerImages()
	}
}


func InitDockerImages() {
	needImages := make(map[string]bool)

	for _, item := range config.CompileObject {
		if item.BuildImage != "-" {
			needImages[item.BuildImage] = true
		}

		if item.RunImage != "-" {
			needImages[item.RunImage] = true
		}
	}

	fmt.Println("[Docker] fetching image list")
	images, err := Client.ImageList(Context, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("[Docker] comparing image list")
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
			BuildImage(imageTag)
		} else {
			PullImage(imageTag)
		}
	}
}