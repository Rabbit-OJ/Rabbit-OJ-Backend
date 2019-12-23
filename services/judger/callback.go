package judger

import (
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/mq"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
)

var (
	CallbackWaitGroup sync.WaitGroup
)

func callbackAllError(status, sid string, storage *Storage) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%s) Callback judge error with status: %s \n", sid, status)

		ceResult := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)

		for i := range ceResult {
			ceResult[i] = &protobuf.JudgeCaseResult{
				Status: status,
			}
		}

		response := &protobuf.JudgeResponse{
			Sid:    sid,
			Result: ceResult,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := mq.Publish(
			config.DefaultExchangeName,
			config.JudgeResultRoutingKey,
			pro); err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func callbackSuccess(sid string, resultList []*protobuf.JudgeCaseResult) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%s) Callback judge success \n", sid)

		response := &protobuf.JudgeResponse{
			Sid:    sid,
			Result: resultList,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := mq.Publish(
			config.DefaultExchangeName,
			config.JudgeResultRoutingKey,
			pro); err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func callbackWebSocket(sid string) {
	websocket.SocketHub.JudgeHub.Broadcast <- sid
}
