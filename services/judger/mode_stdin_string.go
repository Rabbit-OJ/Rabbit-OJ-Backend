package judger

import "fmt"

func ModeStdinString(src, dest string) (bool, uint32) {
	accepted := true
	acceptedCount := uint32(0)

	for {
		rightAnswer, judgeAnswer := "", ""
		if _, err := fmt.Sscanf(src, "%s", &rightAnswer); err != nil {
			break
		}

		if _, err := fmt.Sscanf(dest, "%s", &judgeAnswer); err != nil {
			accepted = false
			break
		}

		if rightAnswer != judgeAnswer {
			accepted = false
			break
		}

		acceptedCount++
	}

	return accepted, acceptedCount
}
