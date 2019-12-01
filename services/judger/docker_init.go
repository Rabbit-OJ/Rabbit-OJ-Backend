package judger

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
)

var (
	DockerContext context.Context
	DockerClient  *client.Client
)

func DockerInit() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	DockerContext, DockerClient = ctx, cli
	if os.Getenv("ENV") == "production" {
		InitImages()
	}
}
