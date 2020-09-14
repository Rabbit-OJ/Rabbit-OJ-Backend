package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/submission"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
)

func JudgeRequestBridge(body []byte) {
	judgeRequest := &protobuf.JudgeRequest{}
	if err := proto.Unmarshal(body, judgeRequest); err != nil {
		fmt.Println(err)

		Requeue(config.JudgeRequestTopicName, body)
		return
	}

	if config.Global.Extensions.Expire.Enabled &&
		judgeRequest.Time-time.Now().Unix() > config.Global.Extensions.CheckJudge.Interval*int64(time.Minute) {
		fmt.Printf("[Bridge] Received expired judge %d , will ignore this\n", judgeRequest.Sid)
		return
	}

	if alreadyAcked, err := Scheduler(judgeRequest); err != nil {
		if !alreadyAcked {
			Requeue(config.JudgeRequestTopicName, body)
		}

		fmt.Println(err)
		return
	}
}

func JudgeResponseBridge(body []byte) {
	judgeResult := &protobuf.JudgeResponse{}
	if err := proto.Unmarshal(body, judgeResult); err != nil {
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
}

func Requeue(topic string, body []byte) {

}