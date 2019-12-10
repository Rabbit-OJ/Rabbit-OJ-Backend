package docker

import (
	"context"
	"github.com/docker/docker/client"
)

var (
	DockerContext context.Context
	DockerClient  *client.Client
)
