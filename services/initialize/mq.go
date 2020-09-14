package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/mq"
	"context"
	"os"
)

func MQ(ctx context.Context) {
	mq.InitKafka(ctx)

	if os.Getenv("Role") == "Judge" {
		judger.JudgeRequestDeliveryChan = make(chan []byte)
		judger.JudgeRequeueDeliveryChan = make(chan []byte)

		mq.JudgeRequestDeliveryChan = judger.JudgeRequestDeliveryChan
		mq.JudgeRequeueDeliveryChan = judger.JudgeRequeueDeliveryChan

		mq.CreateJudgeRequestConsumer([]string{config.JudgeRequestTopicName}, "req1")
		go judger.JudgeRequestHandler()
		go mq.RequeueHandler()
	}

	if os.Getenv("Role") == "Server" {
		judger.JudgeResponseDeliveryChan = make(chan []byte)
		mq.JudgeResponseDeliveryChan = judger.JudgeResponseDeliveryChan

		mq.CreateJudgeResponseConsumer([]string{config.JudgeResponseTopicName}, "res1")
		go judger.JudgeResultHandler()
	}
}
