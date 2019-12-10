package mq

import (
	"github.com/streadway/amqp"
)

var (
	Connection      *amqp.Connection
	ConsumerChannel *amqp.Channel
	PublishChannel  *amqp.Channel
)

func DeclareConsumer(queueName, consumerTag string) (<-chan amqp.Delivery, error) {
	deliveries, err := ConsumerChannel.Consume(
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
	if err := ConsumerChannel.ExchangeDeclare(
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
	_, err := ConsumerChannel.QueueDeclare(
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
	if err := ConsumerChannel.QueueBind(queueName,
		routingKey,
		sourceExchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
