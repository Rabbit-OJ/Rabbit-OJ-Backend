package judger

import (
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func JudgeStart(delivery *amqp.Delivery) {
	judgeRequest := &protobuf.JudgeRequest{}
	if err := proto.Unmarshal(delivery.Body, judgeRequest); err != nil {
		fmt.Println(err)
		return
	}

	if err := Scheduler(judgeRequest); err != nil {
		fmt.Println(err)
		return
	}
}

func JudgeResultStart(delivery *amqp.Delivery) {
	judgeResult := &protobuf.JudgeResponse{}
	if err := proto.Unmarshal(delivery.Body, judgeResult); err != nil {
		fmt.Println(err)
		return
	}

	if err := submission.Result(judgeResult); err != nil {
		fmt.Println(err)
		return
	}
}

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

	return Publish(
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

	return Publish(
		utils.DefaultExchangeName,
		utils.JudgeResultRoutingKey,
		pro)
}

func callbackWebSocket(sid string) {
	websocket.SocketHub.Broadcast <- sid
}
