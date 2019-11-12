package judger

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var (
	DockerContext context.Context
	DockerClient  *client.Client
)

func InitDocker() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	DockerContext, DockerClient = ctx, cli
}
