package judger

import (
	"fmt"
	"math"
)

func ModeStdinFloat64(src, dest string) (bool, uint32) {
	accepted := true
	acceptedCount := uint32(0)

	for {
		var rightAnswer, judgeAnswer float64

		if _, err := fmt.Sscanf(src, "%f", &rightAnswer); err != nil {
			break
		}

		if _, err := fmt.Sscanf(dest, "%f", &judgeAnswer); err != nil {
			accepted = false
			break
		}

		if math.Abs(rightAnswer-judgeAnswer) > 1e-6 {
			accepted = false
			break
		}

		acceptedCount++
	}

	return accepted, acceptedCount
}
