package utils

import "fmt"

func CaseFilePath(tid string, caseId uint32, mode string) string {
	return fmt.Sprintf("./files/judge/%s/%d.%s", tid, caseId, mode)
}

func AvatarPath(uid string) string {
	return fmt.Sprintf("./files/avatar/%s.avatar", uid)
}

func DefaultAvatarPath() string {
	return "./statics/avatar.png"
}