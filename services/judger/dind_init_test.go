package judger

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"testing"
)

func TestPullImage(t *testing.T) {
	DockerInit()
	tag := "node:alpine"

	fmt.Println("[DIND] pulling image : " + tag)
	out, err := DockerClient.ImagePull(DockerContext, tag, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer func() { _ = out.Close() }()

	if _, err := io.Copy(os.Stderr, out); err != nil {
		fmt.Println(err)
	}
}

func TestBuildImage(t *testing.T) {
	DockerInit()
	BuildImage("alpine_tester")
}