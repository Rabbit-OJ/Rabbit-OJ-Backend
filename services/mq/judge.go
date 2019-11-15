package mq

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/submission"
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

	if err := judger.Scheduler(judgeRequest); err != nil {
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
