package mq

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/services/submission"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func JudgeStart(delivery *amqp.Delivery) {

}

func JudgeResultStart(delivery *amqp.Delivery) {
	judgeResult := &protobuf.JudgeResponse{}
	err := proto.Unmarshal(delivery.Body, judgeResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	submissionDetail, err := submission.Detail(judgeResult.Sid)
	if err != nil {
		fmt.Println(err)
		return
	}

	if submissionDetail.Status != "ING" {
		fmt.Printf("Invalid Status" + judgeResult.Sid)
		return
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
		fmt.Println(err)
		return
	}

	if err := submission.Update(judgeResult.Sid, timeUsed, spaceUsed, status, caseJson); err != nil {
		fmt.Println(err)
	}

	if status == "AC" {
		go question.UpdateAcceptedCount(submissionDetail.Tid)
	}
}
