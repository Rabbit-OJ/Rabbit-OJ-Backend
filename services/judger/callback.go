package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/mq"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
	"xorm.io/xorm"
)

var (
	CallbackWaitGroup sync.WaitGroup
)

func callbackAllError(status string, sid uint32, isContest bool, storage *Storage) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%d) Callback judge error with status: %s \n", sid, status)

		ceResult := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)
		for i := range ceResult {
			ceResult[i] = &protobuf.JudgeCaseResult{
				Status: status,
			}
		}

		response := &protobuf.JudgeResponse{
			Sid:       sid,
			Result:    ceResult,
			IsContest: isContest,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := mq.Publish(
			config.DefaultExchangeName,
			config.JudgeResultRoutingKey,
			pro); err != nil {

			fmt.Println(err)
			return
		}
	}()
}

func callbackSuccess(sid uint32, isContest bool, resultList []*protobuf.JudgeCaseResult) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%d) Callback judge success \n", sid)

		response := &protobuf.JudgeResponse{
			Sid:       sid,
			Result:    resultList,
			IsContest: isContest,
		}

		pro, err := proto.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := mq.Publish(
			config.DefaultExchangeName,
			config.JudgeResultRoutingKey,
			pro); err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func callbackWebSocket(sid uint32) {
	judgeHub.Broadcast <- sid
}

func callbackContest(sid uint32, isAccepted bool) {
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
