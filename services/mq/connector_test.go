package mq

import (
	"Rabbit-OJ-Backend/services/config"
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	config.Init()
	Init()
	if err := DeclareExchange("test_exchange", "direct"); err != nil {
		panic(err)
	}
	if err := DeclareQueue("test_queue"); err != nil {
		panic(err)
	}
	if err := BindQueue("test_queue", "test", "test_exchange"); err != nil {
		panic(err)
	}

	done := make(chan bool)
	deliveries, err := DeclareConsumer("test_queue", "")
	if err != nil {
		panic(err)
	}

	go func() {
		for delivery := range deliveries {
			fmt.Println(delivery.Body)

			if err := delivery.Ack(false); err != nil {
				fmt.Println(err)
			}
		}
	}()

	go func() {
		for {
			fmt.Println("pushed")
			if err := Publish("test_exchange", "test", []byte("abc")); err != nil {
				fmt.Println(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	<-done
}
