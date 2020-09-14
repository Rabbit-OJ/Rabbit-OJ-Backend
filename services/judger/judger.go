package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/judger/compare"
	"Rabbit-OJ-Backend/services/tester"
)

func JudgeOneCase(testResult *models.TestResult, stdout, rightStdout, compMode string) *models.JudgeResult {
	result := &models.JudgeResult{}

	if testResult.Status != tester.StatusOK {
		result.Status = testResult.Status
	} else {
		isAC := false
		if compMode == "STDIN_F" {
			isAC, _ = compare.ModeStdinFloat64(stdout, rightStdout)
		} else if compMode == "STDIN_S" {
			isAC, _ = compare.ModeStdinString(stdout, rightStdout)
		} else {
			isAC = compare.ModeCMP(stdout, rightStdout)
		}

		if isAC {
			result.Status = "AC"
		} else {
			result.Status = "WA"
		}
	}

	result.TimeUsed, result.SpaceUsed = testResult.TimeUsed, testResult.SpaceUsed
	return result
}
