package submission

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"encoding/json"
)

func Result(judgeResult *protobuf.JudgeResponse) error {
	submissionDetail, err := Detail(judgeResult.Sid)
	if err != nil {
		return err
	}

	if submissionDetail.Status != "ING" {
		return err
	}

	status, spaceUsed, timeUsed := "AC", uint32(0), uint32(0)
	for _, res := range judgeResult.Result {
		if res.Status != "AC" {
			status = "NO"
			break
		} else {
			spaceUsed += res.SpaceUsed
			timeUsed += res.TimeUsed
		}
	}

	if caseLen := len(judgeResult.Result); caseLen >= 1 {
		timeUsed /= uint32(caseLen)
		spaceUsed /= uint32(caseLen)
	}

	caseJson, err := json.Marshal(judgeResult.Result)
	if err != nil {
		return err
	}

	if err := Update(judgeResult.Sid, timeUsed, spaceUsed, status, caseJson); err != nil {
		return err
	}

	if status == "AC" {
		go question.UpdateAcceptedCount(submissionDetail.Tid)
	}

	return nil
}
