package judger

import (
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/docker/docker/api/types"
	"testing"
)

func TestRunner(t *testing.T) {
	toBeWrite := "test123"

	DockerInit()
	io, err := utils.ConvertToTar([]utils.TarFileBasicInfo{{"test.txt", []byte(toBeWrite)}})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if err := DockerClient.CopyToContainer(
		DockerContext,
		"b202d119039b",
		"/root",
		io,
		types.CopyToContainerOptions{
			AllowOverwriteDirWithFile: true,
			CopyUIDGID:                true,
		}); err != nil {

		fmt.Println(err)
		t.Fail()
		return
	}
}
