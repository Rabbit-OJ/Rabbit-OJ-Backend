package compare

import (
	"math"
	"strconv"
	"strings"
)

func ModeStdinFloat64(src, dest string) (bool, uint32) {
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

		srcFloat, srcErr := strconv.ParseFloat(srcArr[i], 64)
		destFloat, destErr := strconv.ParseFloat(destArr[i], 64)

		if srcErr != nil || destErr != nil {
			accepted = false
			continue
		}

		if math.Abs(srcFloat - destFloat) <= 1e-6 {
			acceptedCount++
		} else {
			accepted = false
		}
	}

	return accepted, acceptedCount
}
