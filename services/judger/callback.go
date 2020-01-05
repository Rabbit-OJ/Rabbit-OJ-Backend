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
)

var (
	CallbackWaitGroup sync.WaitGroup
)

func callbackAllError(status, sid string, isContest bool, storage *Storage) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%s) Callback judge error with status: %s \n", sid, status)

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

func callbackSuccess(sid string, isContest bool, resultList []*protobuf.JudgeCaseResult) {
	go func() {
		CallbackWaitGroup.Add(1)
		defer CallbackWaitGroup.Done()

		fmt.Printf("(%s) Callback judge success \n", sid)

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

func callbackWebSocket(sid string, isContest, accepted bool) {
	judgeHub.Broadcast <- sid

	if isContest {
		callbackContest(sid, accepted)
	}
}

func callbackContest(sid string, isAccepted bool) {
	tx := db.DB.Begin()

	status := contest.StatusPending
	if isAccepted {
		status = contest.StatusAC
	} else {
		status = contest.StatusERR
	}

	if err := contest.ChangeSubmitState(tx, sid, status);
		err != nil {

		fmt.Println(err)
		tx.Rollback()
		return
	}

	submissionInfo, err := contest.SubmissionInfo(tx, sid)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	if err := contest.RegenerateUserScore(tx, submissionInfo.Cid, submissionInfo.Uid);
		err != nil {

		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()
}
