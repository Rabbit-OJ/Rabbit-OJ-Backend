package initialize

import (
	"Rabbit-OJ-Backend/services/channel"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/services/submission"
	"context"
	"os"
)

func MQ(ctx context.Context) {
	mq.InitKafka(ctx)

	channel.MQPublishMessageChannel = make(chan *channel.MQMessage)
	if os.Getenv("Role") == "Judge" {
		channel.JudgeRequestDeliveryChan = make(chan []byte)
		channel.JudgeRequestBridgeChan = make(chan *channel.JudgeRequestBridgeMessage)

		mq.CreateJudgeRequestConsumer([]string{config.JudgeRequestTopicName}, "req1")
		go judger.JudgeRequestHandler()
		go submission.MachineJudgeRequestBridge()
	}

	if os.Getenv("Role") == "Server" {
		channel.JudgeResponseDeliveryChan = make(chan []byte)

		mq.CreateJudgeResponseConsumer([]string{config.JudgeResponseTopicName}, "res1")
		go submission.JudgeResultHandler()
	}
	go mq.PublishService()
}
