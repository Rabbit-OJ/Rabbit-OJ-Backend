package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/submission"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"time"
)

func JudgeRequestBridge(delivery *amqp.Delivery, okChan chan bool) {
	defer func() {
		okChan <- true
	}()

	body := delivery.Body
	judgeRequest := &protobuf.JudgeRequest{}
	if err := proto.Unmarshal(body, judgeRequest); err != nil {
		fmt.Println(err)

		if err := delivery.Nack(false, true); err != nil {
			fmt.Println(err)
		}
		return
	}

	if config.Global.Extensions.Expire.Enabled &&
		judgeRequest.Time-time.Now().Unix() > config.Global.Extensions.CheckJudge.Interval*int64(time.Minute) {
		fmt.Printf("[Bridge] Received expired judge %d , will ignore this\n", judgeRequest.Sid)

		if err := delivery.Ack(false); err != nil {
			fmt.Println(err)
		}
		return
	}

	if alreadyAcked, err := Scheduler(delivery, judgeRequest); err != nil {
		if !alreadyAcked {
			if err := delivery.Nack(false, true); err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println(err)
		return
	}
}

func JudgeResponseBridge(delivery *amqp.Delivery) {
	judgeResult := &protobuf.JudgeResponse{}
	if err := proto.Unmarshal(delivery.Body, judgeResult); err != nil {
		fmt.Println(err)
		return
	}

	status, err := submission.Result(judgeResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	if judgeResult.IsContest {
		callbackContest(judgeResult.Sid, status == "AC")
	}
	go callbackWebSocket(judgeResult.Sid)

	if err := delivery.Ack(false); err != nil {
		fmt.Println(err)
	}
}
