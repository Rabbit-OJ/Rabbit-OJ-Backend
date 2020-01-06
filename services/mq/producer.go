package mq

import (
	"github.com/streadway/amqp"
)

func Publish(exchangeName, routingKey string, body []byte) error {
	//if err := PublishChannel.Confirm(false); err != nil {
	//	fmt.Println(err)
	//
	//	return err
	//} else {
	//	//confirms := PublishChannel.NotifyPublish(make(chan amqp.Confirmation, 1))
	//
	//	//defer func() {
	//	//	if confirmed := <-confirms; confirmed.Ack {
	//	//		fmt.Printf("[MQ] Confirmed Consumer with tag: %d \n", confirmed.DeliveryTag)
	//	//	} else {
	//	//		fmt.Printf("[MQ] Failed Confirmed Consumer with tag: %d \n", confirmed.DeliveryTag)
	//	//	}
	//	//}()
	//}

	if err := PublishChannel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Body: body,
		}); err != nil {
		return err
	}

	return nil
}
