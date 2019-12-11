package files

import (
	"archive/tar"
	"bytes"
	"fmt"
)

// this function is mainly modified from https://golang.org/pkg/archive/tar/

type TarFileBasicInfo struct {
	Name string
	Body []byte
}

func ConvertToTar(files []TarFileBasicInfo) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := tar.NewWriter(&buf)

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0644,
			Size: int64(len(file.Body)),
		}

		if err := writer.WriteHeader(hdr); err != nil {
			fmt.Println(err)
		}
		if _, err := writer.Write(file.Body); err != nil {
			fmt.Println(err)
		}
	}

	if err := writer.Close(); err != nil {
		fmt.Println(err)
	}
	return &buf, nil
}
