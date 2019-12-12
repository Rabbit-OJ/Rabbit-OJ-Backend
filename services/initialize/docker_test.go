package initialize

import (
	"Rabbit-OJ-Backend/services/docker"
	"testing"
)

func TestDocker(t *testing.T) {
	Docker()
	docker.BuildImage("alpine_tester:latest")
}
