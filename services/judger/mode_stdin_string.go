package judger

import (
	"strings"
)

func ModeStdinString(src, dest string) (bool, uint32) {
	accepted, acceptedCount := true, uint32(0)

	srcArr, destArr := strings.Fields(src), strings.Fields(dest)
	if len(srcArr) != len(destArr) {
		accepted = false
	}

	for i := 0; i < len(srcArr); i++ {
		if i >= len(destArr) {
			accepted = false
			break
		}

		if srcArr[i] == destArr[i] {
			acceptedCount++
		} else {
			accepted = false
		}
	}

	return accepted, acceptedCount
}
