package mq

import (
	"context"
	"github.com/Shopify/sarama"
)

var (
	Version   sarama.KafkaVersion
	CancelCtx context.Context

	JudgeRequestDeliveryChan  chan []byte
	JudgeRequeueDeliveryChan  chan []byte
	JudgeResponseDeliveryChan chan []byte
)

func InitKafka(ctx context.Context) {
	if version, err := sarama.ParseKafkaVersion("2.6.0"); err != nil {
		panic(err)
	} else {
		Version = version
	}
	CancelCtx = ctx
	InitProducer()
}
