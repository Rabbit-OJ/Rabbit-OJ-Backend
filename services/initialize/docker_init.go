package initialize

import (
	"Rabbit-OJ-Backend/services/docker"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
)

func Docker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	docker.Context, docker.Client = ctx, cli
	if os.Getenv("ENV") == "production" {
		DockerImages()
	}
}
