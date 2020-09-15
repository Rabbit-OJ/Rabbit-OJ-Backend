package submission

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/channel"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/judger"
	"fmt"
	"time"
	"xorm.io/xorm"

	"github.com/golang/protobuf/proto"
)

func JudgeRequestBridge(data *channel.JudgeRequestBridgeMessage) {
	body := data.Data
	defer func() {
		data.SuccessChan <- true
	}()

	judgeRequest := &protobuf.JudgeRequest{}
	if err := proto.Unmarshal(body, judgeRequest); err != nil {
		fmt.Println(err)
		return
	}

	if config.Global.Extensions.Expire.Enabled &&
		judgeRequest.Time-time.Now().Unix() > config.Global.Extensions.CheckJudge.Interval*int64(time.Minute) {
		fmt.Printf("[Bridge] Received expired judge %d , will ignore this\n", judgeRequest.Sid)
		return
	}

	if alreadyAcked, err := judger.Scheduler(judgeRequest); err != nil {
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

	status, err := Result(judgeResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	if judgeResult.IsContest {
		CallbackContest(judgeResult.Sid, status == "AC")
	}
	go CallbackWebSocket(judgeResult.Sid)
}

func Requeue(topic string, body []byte) {
	// todo
}


func CallbackWebSocket(sid uint32) {
	JudgeHub.Broadcast <- sid
}

func CallbackContest(sid uint32, isAccepted bool) {
	_, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		status := contest.StatusPending
		if isAccepted {
			status = contest.StatusAC
		} else {
			status = contest.StatusERR
		}

		if err := contest.ChangeSubmitState(session, sid, status); err != nil {
			return nil, err
		}

		submissionInfo, err := contest.SubmissionInfo(session, sid)
		if err != nil {
			return nil, err
		}

		if err := contest.RegenerateUserScore(session,
			submissionInfo.Cid, submissionInfo.Uid,
			isAccepted); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func JudgeResultHandler() {
	for delivery := range channel.JudgeResponseDeliveryChan {
		go JudgeResponseBridge(delivery)
	}
}

func MachineJudgeRequestBridge() {
	for delivery := range channel.JudgeRequestBridgeChan {
		go JudgeRequestBridge(delivery)
	}
}