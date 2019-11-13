package judger

import (
	"github.com/docker/docker/api/types"
)

func ShouldDockerPullImage() (bool, error) {
	images, err := DockerClient.ImageList(DockerContext, types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	answer := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == "rabbit_oj_tester:latest" {
				answer = true
				break
			}
		}

		if answer {
			break
		}
	}

	return answer, nil
}

func DockerBuildImage() {

}
