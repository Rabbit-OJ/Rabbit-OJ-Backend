package docker

import (
	"context"
	"github.com/docker/docker/client"
)

var (
	Context context.Context
	Client  *client.Client
)
