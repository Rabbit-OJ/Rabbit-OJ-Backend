package utils

import (
	"fmt"
	"os"
	"time"
)

func CaseFilePath(tid string, caseId uint32, mode string) string {
	return fmt.Sprintf("./files/judge/%s/%d.%s", tid, caseId, mode)
}

func AvatarPath(uid string) string {
	return fmt.Sprintf("./files/avatar/%s.avatar", uid)
}

func DefaultAvatarPath() string {
	return "./statics/avatar.png"
}

func CodeDir() string {
	return "./files/submission/"
}

func CodePath(uuid string) string {
	return fmt.Sprintf("./files/submission/%s.code", uuid)
}

func CodeGenerateFileName(uid string) (string, error) {
	t := time.Now()
	path := CodeDir() + t.Format("2006/01/02")

	if !Exists(path) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s/%s_%d", t.Format("2006/01/02"), uid, t.UnixNano()), nil
}
