package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

func callbackAllError(status, sid string, storage *Storage) error {
	fmt.Println("Callback judge error : " + sid + " with status " + status)

	ceResult := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)

	for i := range ceResult {
		ceResult[i].Status = status
	}

	response := &protobuf.JudgeResponse{
		Sid:    sid,
		Result: ceResult,
	}

	pro, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	return mq.Publish(
		utils.DefaultExchangeName,
		utils.JudgeResultRoutingKey,
		pro)
}

func callbackSuccess(sid string, resultList []*protobuf.JudgeCaseResult) error {
	fmt.Println("Callback judge success : " + sid)

	response := &protobuf.JudgeResponse{
		Sid:    sid,
		Result: resultList,
	}

	pro, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	return mq.Publish(
		utils.DefaultExchangeName,
		utils.JudgeResultRoutingKey,
		pro)
}