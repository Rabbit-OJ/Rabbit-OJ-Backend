package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/mq"
	"context"
	"os"
)

func MQ(ctx context.Context) {
	if os.Getenv("Role") == "Judge" {
		judger.JudgeRequestDeliveryChan = make(chan []byte)
		judger.JudgeRequeueDeliveryChan = make(chan []byte)

		mq.JudgeRequestDeliveryChan = judger.JudgeRequestDeliveryChan
		mq.JudgeRequeueDeliveryChan = judger.JudgeRequeueDeliveryChan

		go judger.JudgeRequestHandler()
		go mq.RequeueHandler()
		mq.CreateJudgeRequestConsumer([]string{config.JudgeRequestTopicName}, "req1")
	}

	if os.Getenv("Role") == "Server" {
		judger.JudgeResponseDeliveryChan = make(chan []byte)
		mq.JudgeRequestDeliveryChan = judger.JudgeResponseDeliveryChan

		go judger.JudgeResultHandler()
		mq.CreateJudgeResponseConsumer([]string{config.JudgeResponseTopicName}, "res1")
	}

	mq.InitKafka(ctx)
}
// TODO: handle Reconnect
//func handleReconnect(closeChan chan *amqp.Error) {
//	select {
//	case err := <-closeChan:
//		fmt.Printf("Reconnecting rabbitmq, meet error: %+v \n", err)
//		connect()
//	}
//}
