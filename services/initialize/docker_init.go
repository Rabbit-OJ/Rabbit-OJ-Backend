package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/docker"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func Docker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	docker.Context, docker.Client = ctx, cli
	if config.Global.Extensions.AutoPull {
		DockerImages()
	}
}
