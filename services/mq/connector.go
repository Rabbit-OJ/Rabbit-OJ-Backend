package mq

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"github.com/streadway/amqp"
)

var (
	Connection      *amqp.Connection
	Channel         *amqp.Channel
)

func handleReconnect(closeChan chan *amqp.Error) {
	select {
	case err := <-closeChan:
		fmt.Printf("Reconnecting rabbitmq, meet error: %+v \n", err)
		Init()
		// todo: declare consumer
	}
}

func Init() {
	connStr := config.Global.RabbitMQ
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

	closeChan := make(chan *amqp.Error)
	Connection.NotifyClose(closeChan)
	go handleReconnect(closeChan)
}

func DeclareConsumer(queueName, consumerTag string) (<-chan amqp.Delivery, error) {
	deliveries, err := Channel.Consume(
		queueName,
		consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return deliveries, err
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
