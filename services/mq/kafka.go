package mq

import (
	"context"
	"os"

	"github.com/Shopify/sarama"
)

var (
	Version               sarama.KafkaVersion
	Broker                []string
	CancelCtx, CancelFunc = context.WithCancel(context.Background())

	JudgeRequestDeliveryChan  chan []byte
	JudgeResponseDeliveryChan chan []byte
)

func InitKafka() {
	devEnv := os.Getenv("DEV")
	if devEnv == "1" {
		Broker = []string{
			"localhost:9092",
			"localhost:9093",
			"localhost:9094",
		}
	} else {
		Broker = []string{
			"kafka:9092",
		}
	}

	if version, err := sarama.ParseKafkaVersion("2.6.0"); err != nil {
		panic(err)
	} else {
		Version = version
	}

	Producer()
}
