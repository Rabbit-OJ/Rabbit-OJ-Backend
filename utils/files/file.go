package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	MagicBytes = []byte{'\x7F', 'E', 'L', 'F'}
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func TouchFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write([]byte(""))
	return err
}

func TouchFileWithMagic(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(MagicBytes)
	return err
}

func ReadFileBytes(absPath string) ([]byte, error) {
	path, err := filepath.Abs(absPath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(path)
}