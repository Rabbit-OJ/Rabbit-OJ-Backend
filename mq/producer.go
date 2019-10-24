package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Publish(exchangeName, routingKey string, body []byte) error {
	if err := Channel.Confirm(false); err != nil {
		fmt.Println(err)
	} else {
		confirms := Channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer func() {
			if confirmed := <-confirms; confirmed.Ack {
				fmt.Printf("[MQ] Confirmed Consumer with tag: %d \n", confirmed.DeliveryTag)
			} else {
				fmt.Printf("[MQ] Failed Confirmed Consumer with tag: %d \n", confirmed.DeliveryTag)
			}
		}()
	}

	if err := Channel.Publish(exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body:         body,
			DeliveryMode: amqp.Transient,
			Priority:     0,
		}); err != nil {
		return err
	}

	return nil
}
