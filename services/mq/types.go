package mq

import (
	"github.com/Shopify/sarama"
)

type JudgeRequestConsumer struct {
	ready chan bool
}

func (consumer *JudgeRequestConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *JudgeRequestConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *JudgeRequestConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		go func() {
			JudgeRequestDeliveryChan <- message.Value
		}()
		session.MarkMessage(message, "")
	}

	return nil
}

type JudgeResponseConsumer struct {
	ready chan bool
}

func (consumer *JudgeResponseConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *JudgeResponseConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *JudgeResponseConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		go func() {
			JudgeResponseDeliveryChan <- message.Value
		}()
		session.MarkMessage(message, "")
	}

	return nil
}
