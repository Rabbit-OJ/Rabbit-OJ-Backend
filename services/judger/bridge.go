package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/submission"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func JudgeRequestBridge(delivery *amqp.Delivery, okChan chan bool) {
	defer func() {
		okChan <- true
	}()

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

func JudgeResponseBridge(delivery *amqp.Delivery, okChan chan bool) {
	defer func() {
		okChan <- true
	}()

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
