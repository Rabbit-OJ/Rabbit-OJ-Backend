package utils

import (
	"archive/tar"
	"bytes"
	"fmt"
	"log"
)

// this function is mainly modified from https://golang.org/pkg/archive/tar/

type tarFileBasicInfo struct {
	Name string
	Body []byte
}

func ConvertToTar(fileName string, code []byte) *tar.Reader {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	defer func() {
		err := tw.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	files := []tarFileBasicInfo{
		{fileName, code},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0644,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write(file.Body); err != nil {
			log.Fatal(err)
		}
	}

	return tar.NewReader(&buf)
}
