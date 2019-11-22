package judger

import (
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/utils"
	"os"
)

func InitMQ() {
	if err := mq.DeclareExchange(utils.DefaultExchangeName, "direct"); err != nil {
		panic(err)
	}

	if err := mq.DeclareQueue(utils.JudgeQueueName); err != nil {
		panic(err)
	}

	if err := mq.DeclareQueue(utils.JudgeResultQueueName); err != nil {
		panic(err)
	}

	if err := mq.BindQueue(utils.JudgeQueueName, utils.JudgeRoutingKey, utils.DefaultExchangeName); err != nil {
		panic(err)
	}

	if err := mq.BindQueue(utils.JudgeResultQueueName, utils.JudgeResultRoutingKey, utils.DefaultExchangeName); err != nil {
		panic(err)
	}

	if os.Getenv("Role") == "Judge" {
		// judge mode
		deliveries, err := mq.DeclareConsumer(utils.JudgeQueueName, utils.JudgeRoutingKey)
		if err != nil {
			panic(err)
		}

		go JudgeHandler(deliveries)
	}

	if os.Getenv("Role") == "Server" {
		// server mode
		deliveries, err := mq.DeclareConsumer(utils.JudgeResultQueueName, utils.JudgeResultRoutingKey)
		if err != nil {
			panic(err)
		}

		go JudgeResultHandler(deliveries)
	}
}
