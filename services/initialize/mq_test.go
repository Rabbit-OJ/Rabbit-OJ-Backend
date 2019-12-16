package initialize

import (
	"Rabbit-OJ-Backend/services/mq"
	"fmt"
	"testing"
	"time"
)

func TestMQ(t *testing.T) {
	exitChan := make(chan bool)
	Config()
	MQ(exitChan)

	if err := mq.DeclareExchange("test_exchange", "direct"); err != nil {
		panic(err)
	}
	if err := mq.DeclareQueue("test_queue"); err != nil {
		panic(err)
	}
	if err := mq.BindQueue("test_queue", "test", "test_exchange"); err != nil {
		panic(err)
	}
	go publisher()
	<-exitChan
}

func publisher() {
	cnt := 0
	for {
		fmt.Println("[Publisher] Enqueue a message to the queue")
		cnt++

		if err := mq.Publish(
			"test_exchange",
			"test",
			[]byte{byte(cnt)}); err != nil {

			fmt.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}
