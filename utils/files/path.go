package files

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func AvatarPath(uid uint32) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/avatar/%d.avatar", uid))
}

func DefaultAvatarPath() (string, error) {
	return filepath.Abs("./files/statics/avatar.png")
}

func CodeDir() string {
	return "./files/submission/"
}

func CodePath(uuid string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/submission/%s.code", uuid))
}

func CodeGenerateFileNameWithMkdir(uid uint32) (string, error) {
	t := time.Now()
	path := CodeDir() + t.Format("2006/01/02")

	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s/%d_%d", t.Format("2006/01/02"), uid, t.UnixNano()), nil
}

func StorageFilePath() (string, error) {
	return filepath.Abs("./files/storage.json")
}

func ConfigFilePath() (string, error) {
	return filepath.Abs("./config.json")
}

func JudgeDirPathWithMkdir(tid uint32, version string) (string, error) {
	path, err := filepath.Abs(fmt.Sprintf("./files/judge/%d/%s", tid, version))
	if err != nil {
		return "", err
	}

	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return "", err
		}
	}
	return path, nil
}

func JudgeFilePath(tid uint32, version, caseId, caseType string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/judge/%d/%s/%s.%s", tid, version, caseId, caseType))
}

func CertFilePath(fileName string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/certs/%s", fileName))
}
