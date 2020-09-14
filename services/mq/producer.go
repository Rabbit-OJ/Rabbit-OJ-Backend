package mq

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/Shopify/sarama"
)

var (
	SyncProducer  sarama.SyncProducer
	AsyncProducer sarama.AsyncProducer
)

func InitSyncProducer() {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Errors = true
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.NoResponse

	producer, err := sarama.NewSyncProducer(config.Global.Kafka.Brokers, saramaConfig)
	if err != nil {
		panic(err)
	}

	SyncProducer = producer
	<-CancelCtx.Done()
	_ = SyncProducer.Close()
}

func InitAsyncProducer() {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Retry.Max = 5
	saramaConfig.Producer.Return.Errors = false

	producer, err := sarama.NewAsyncProducer(config.Global.Kafka.Brokers, saramaConfig)
	if err != nil {
		panic(err)
	}

	AsyncProducer = producer
	<-CancelCtx.Done()
	_ = AsyncProducer.Close()
}

func InitProducer() {
	go InitSyncProducer()
	go InitAsyncProducer()
}

func PublishMessage(topic string, key, buf []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(buf),
	}

	_, _, err := SyncProducer.SendMessage(msg)
	return err
}
