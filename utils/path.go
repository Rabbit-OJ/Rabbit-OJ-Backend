package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func AvatarPath(uid string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/avatar/%s.avatar", uid))
}

func DefaultAvatarPath() (string, error) {
	return filepath.Abs("./statics/avatar.png")
}

func CodeDir() string {
	return "./files/submission/"
}

func CodePath(uuid string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/submission/%s.code", uuid))
}

func CodeGenerateFileNameWithMkdir(uid string) (string, error) {
	t := time.Now()
	path := CodeDir() + t.Format("2006/01/02")

	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s/%s_%d", t.Format("2006/01/02"), uid, t.UnixNano()), nil
}

func DockerCasePath(caseId int64) string {
	return fmt.Sprintf("/case/%d.in", caseId)
}

func DockerOutputPath(caseId int64) string {
	return fmt.Sprintf("/output/%d.in", caseId)
}

func StorageFilePath() (string, error) {
	return filepath.Abs("./files/storage.json")
}

func JudgeDirPathWithMkdir(tid, version string) (string, error) {
	path, err :=  filepath.Abs(fmt.Sprintf("./files/judge/%s/%s", tid, version))
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

func JudgeFilePath(tid, version, caseId, caseType string) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/judge/%s/%s/%s.%s", tid, version, caseId, caseType))
}
