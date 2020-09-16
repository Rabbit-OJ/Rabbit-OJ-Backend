package submission

import (
	"Rabbit-OJ-Backend/models"
	JudgerModels "Rabbit-OJ-Backend/services/judger/models"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/services/user"
)

func Result(sid uint32, judgeResult []*JudgerModels.JudgeResult) (string, error) {
	submissionDetail, err := Detail(sid)
	if err != nil {
		return "", err
	}

	if submissionDetail.Status != "ING" {
		return "", err
	}

	status, spaceUsed, timeUsed := "AC", uint32(0), uint32(0)
	for _, res := range judgeResult {
		spaceUsed += res.SpaceUsed
		timeUsed += res.TimeUsed

		if res.Status != "AC" {
			status = "NO"
		}
	}

	if caseLen := len(judgeResult); caseLen >= 1 {
		timeUsed /= uint32(caseLen)
		spaceUsed /= uint32(caseLen)
	}

	resultObj := make([]models.JudgeResult, len(judgeResult))
	for i, item := range judgeResult {
		resultObj[i] = models.JudgeResult{
			Status:    item.Status,
			TimeUsed:  item.TimeUsed,
			SpaceUsed: item.SpaceUsed,
		}
	}

	if err := Update(sid, timeUsed, spaceUsed, status, resultObj); err != nil {
		return "", err
	}

	if status == "AC" {
		go question.UpdateAcceptedCount(submissionDetail.Tid)
		go user.UpdateAcceptedCount(submissionDetail.Uid)
	}

	return status, nil
}
