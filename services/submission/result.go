package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/judger/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/services/user"
)

func Result(judgeResult *protobuf.JudgeResponse) (string, error) {
	submissionDetail, err := Detail(judgeResult.Sid)
	if err != nil {
		return "", err
	}

	if submissionDetail.Status != "ING" {
		return "", err
	}

	status, spaceUsed, timeUsed := "AC", uint32(0), uint32(0)
	for _, res := range judgeResult.Result {
		spaceUsed += res.SpaceUsed
		timeUsed += res.TimeUsed

		if res.Status != "AC" {
			status = "NO"
		}
	}

	if caseLen := len(judgeResult.Result); caseLen >= 1 {
		timeUsed /= uint32(caseLen)
		spaceUsed /= uint32(caseLen)
	}

	resultObj := make([]models.JudgeResult, len(judgeResult.Result))
	for i, item := range judgeResult.Result {
		resultObj[i] = models.JudgeResult{
			Status:    item.Status,
			TimeUsed:  item.TimeUsed,
			SpaceUsed: item.SpaceUsed,
		}
	}

	if err := Update(judgeResult.Sid, timeUsed, spaceUsed, status, resultObj); err != nil {
		return "", err
	}

	if status == "AC" {
		go question.UpdateAcceptedCount(submissionDetail.Tid)
		go user.UpdateAcceptedCount(submissionDetail.Uid)
	}

	return status, nil
}
