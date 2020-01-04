package submission

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/services/user"
	"encoding/json"
)

func Result(judgeResult *protobuf.JudgeResponse) (string, error) {
	submissionDetail, err := Detail(judgeResult.Sid)
	if err != nil {
		return "", err
	}

	if submissionDetail.Status != "ING" {
		return "", err
	}

	status, spaceUsed, timeUsed := "AC", float64(0), uint32(0)
	for _, res := range judgeResult.Result {
		spaceUsed += res.SpaceUsed
		timeUsed += res.TimeUsed

		if res.Status != "AC" {
			status = "NO"
		}
	}

	if caseLen := len(judgeResult.Result); caseLen >= 1 {
		timeUsed /= uint32(caseLen)
		spaceUsed /= float64(caseLen)
	}

	caseJson, err := json.Marshal(judgeResult.Result)
	if err != nil {
		return "", err
	}

	if err := Update(judgeResult.Sid, timeUsed, spaceUsed, status, caseJson); err != nil {
		return "", err
	}

	if status == "AC" {
		go question.UpdateAcceptedCount(submissionDetail.Tid)
		go user.UpdateAcceptedCount(submissionDetail.Uid)
	}

	return status, nil
}
