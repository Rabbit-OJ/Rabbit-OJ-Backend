package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func MQ(exitChan chan bool) {
	connect()

	go func() {
		<-exitChan
		if err := mq.Connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}

func connect() {
	connStr := config.Global.RabbitMQ
	if conn, err := amqp.Dial(connStr); err != nil {
		panic(err)
	} else {
		mq.Connection = conn
	}

	if channel, err := mq.Connection.Channel(); err != nil {
		panic(err)
	} else {
		mq.ConsumerChannel = channel
	}

	if channel, err := mq.Connection.Channel(); err != nil {
		panic(err)
	} else {
		mq.PublishChannel = channel
	}

	declareServices()

	closeChan := make(chan *amqp.Error)
	mq.Connection.NotifyClose(closeChan)
	go handleReconnect(closeChan)
}

func declareServices() {
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

		go judger.JudgeHandler(deliveries)
	}

	if os.Getenv("Role") == "Server" {
		// server mode
		deliveries, err := mq.DeclareConsumer(utils.JudgeResultQueueName, utils.JudgeResultRoutingKey)
		if err != nil {
			panic(err)
		}

		go judger.JudgeResultHandler(deliveries)
	}
}

func handleReconnect(closeChan chan *amqp.Error) {
	select {
	case err := <-closeChan:
		fmt.Printf("Reconnecting rabbitmq, meet error: %+v \n", err)
		connect()
		// todo: re-declare consumer
	}
}
