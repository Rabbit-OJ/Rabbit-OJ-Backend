package mq

import (
	"github.com/Shopify/sarama"
)

var (
	SyncProducer sarama.SyncProducer
)

func Producer() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(Broker, config)
	if err != nil {
		panic(err)
	}

	SyncProducer = producer
}

// TODO: ACK logic
func PublishMessage(topic string, key, buf []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(buf),
	}

	_, _, err := SyncProducer.SendMessage(msg)
	return err
}
