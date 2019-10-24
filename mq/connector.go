package mq

import (
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

var (
	Connection *amqp.Connection
	Channel    *amqp.Channel
)

func Init() {
	username := "user"
	password := "P@ssw0rd"

	server := "localhost:5672"
	if os.Getenv("ENV") == "production" {
		server = "rabbitmq:5672"
	}

	connStr := fmt.Sprintf("amqp://%s:%s@%s/", username, password, server)

	if conn, err := amqp.Dial(connStr); err != nil {
		panic(err)
	} else {
		Connection = conn
	}

	if channel, err := Connection.Channel(); err != nil {
		panic(err)
	} else {
		Channel = channel
	}

	if err := DeclareExchange(utils.DefaultExchangeName, "direct"); err != nil {
		panic(err)
	}

	if err := DeclareQueue(utils.CaseQueueName); err != nil {
		panic(err)
	}

	if err := DeclareQueue(utils.JudgeQueueName); err != nil {
		panic(err)
	}

	if err := BindQueue(utils.JudgeQueueName, utils.JudgeRoutingKey, utils.DefaultExchangeName); err != nil {
		panic(err)
	}

	if err := BindQueue(utils.CaseQueueName, utils.CaseRoutingKey, utils.DefaultExchangeName); err != nil {
		panic(err)
	}
}

func DeclareExchange(exchangeName, exchangeType string) error {
	if err := Channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}

func DeclareQueue(queueName string) error {
	_, err := Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func BindQueue(queueName, routingKey, sourceExchange string) error {
	if err := Channel.QueueBind(queueName,
		routingKey,
		sourceExchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
