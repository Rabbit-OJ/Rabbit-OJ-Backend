package submission

import (
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	JudgerModels "github.com/Rabbit-OJ/Rabbit-OJ-Judger/models"
	"xorm.io/xorm"
)

func JudgeResponseCallback(sid uint32, isContest bool, judgeResult []*JudgerModels.JudgeResult) {
	status, err := Result(sid, judgeResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isContest {
		CallbackContest(sid, status == "AC")
	}
	go CallbackWebSocket(sid)
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
