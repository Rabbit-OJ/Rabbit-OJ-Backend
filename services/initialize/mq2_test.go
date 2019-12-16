package initialize

import (
	"Rabbit-OJ-Backend/services/mq"
	"fmt"
	"testing"
	"time"
)

func TestMQ2(t *testing.T) {
	exitChan := make(chan bool)
	Config()
	MQ(exitChan)

	go consumer()
	<-exitChan
}

func consumer() {
	deliveries, err := mq.DeclareConsumer("test_queue", "test")
	if err != nil {
		panic(err)
	}

	for delivery := range deliveries {
		fmt.Println(delivery.Body)
		time.Sleep(10 * time.Second)
		if err := delivery.Ack(false); err != nil {
			fmt.Println(err)
		}
	}
}