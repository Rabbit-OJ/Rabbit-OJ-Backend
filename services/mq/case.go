package mq

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func TestCaseConsumer(delivery *amqp.Delivery) {
	testCaseRequest := &protobuf.TestCaseRequest{}
	err := proto.Unmarshal(delivery.Body, testCaseRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	tid := testCaseRequest.Tid
	testCase, err := question.Case(tid)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := proto.Marshal(testCase)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Publish(
		utils.DefaultExchangeName,
		utils.CaseRoutingKey,
		body,
	)

	if err != nil {
		fmt.Println(err)
	}
}

func TestCasePublish(tid string) error {
	testCaseInfo, err := question.Case(tid)

	if err != nil {
		return err
	}

	testCaseRequest := &protobuf.TestCaseResponse{
		Tid:     tid,
		Version: testCaseInfo.Version,
	}
	data, err := proto.Marshal(testCaseRequest)
	if err != nil {
		return err
	}

	return Publish(
		utils.DefaultExchangeName,
		utils.CaseRoutingKey,
		data,
	)
}
