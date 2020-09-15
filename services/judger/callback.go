package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/channel"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/storage"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
)

var (
	CallbackWaitGroup sync.WaitGroup
)

func CallbackAllError(status string, sid uint32, isContest bool, storage *storage.Storage) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%d) Callback judge error with status: %s \n", sid, status)

		ceResult := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)
		for i := range ceResult {
			ceResult[i] = &protobuf.JudgeCaseResult{
				Status: status,
			}
		}

		response := &protobuf.JudgeResponse{
			Sid:       sid,
			Result:    ceResult,
			IsContest: isContest,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		channel.MQPublishMessageChannel <- &channel.MQMessage{
			Async: true,
			Topic: []string{config.JudgeResponseTopicName},
			Key:   []byte(fmt.Sprintf("%d", sid)),
			Value: pro,
		}
	}()
}

func CallbackSuccess(sid uint32, isContest bool, resultList []*protobuf.JudgeCaseResult) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%d) Callback judge success \n", sid)

		response := &protobuf.JudgeResponse{
			Sid:       sid,
			Result:    resultList,
			IsContest: isContest,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		channel.MQPublishMessageChannel <- &channel.MQMessage{
			Async: true,
			Topic: []string{config.JudgeResponseTopicName},
			Key:   []byte(fmt.Sprintf("%d", sid)),
			Value: pro,
		}
	}()
}
