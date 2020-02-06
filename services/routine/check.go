package routine

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/utils/files"
	"context"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	Context       context.Context
	CancelContext context.CancelFunc
)

type toBeRejectObject struct {
	IsContest bool
	Sid       uint32
}

func StartCheck() {
	if config.Global.Extensions.CheckJudge.Enabled {
		Context, CancelContext = context.WithCancel(context.Background())
		go checkRoutine(config.Global.Extensions.CheckJudge.Interval, Context)
	}
}

func StopCheck() {
	CancelContext()
}

func checkRoutine(interval int64, ctx context.Context) {
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Minute):
			handleCheck()
		case <-ctx.Done():
			break
		}
	}
}

type questionJudgeMemoType struct {
	judge  *models.QuestionJudge
	detail *models.QuestionDetail
}

func handleCheck() {
	fmt.Printf("[Judge Check] Start routine \n")

	someMinutesBefore := time.
		Now().
		Add(-1 * time.Duration(config.Global.Extensions.CheckJudge.Interval) * time.Minute)

	var timeoutSubmissions []models.Submission
	if err := db.DB.Table("submission").
		Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
		Find(&timeoutSubmissions); err != nil {
		fmt.Println(err)
		return
	}

	if config.Global.Extensions.CheckJudge.Requeue {
		toBeRejected, questionMemo := make([]*toBeRejectObject, 0), make(map[uint32]questionJudgeMemoType)
		for _, item := range timeoutSubmissions {
			path, err := files.CodePath(item.FileName)
			if err != nil {
				fmt.Println(err)
				toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
				continue
			}
			exist := files.Exists(path)
			if !exist {
				toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
				continue
			}
			code, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
				continue
			}

			if _, ok := questionMemo[item.Tid]; !ok {
				questionDetail, err := question.Detail(item.Tid)
				if err != nil {
					toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
					fmt.Println(err)
					continue
				}
				questionJudge, err := question.JudgeInfo(item.Tid)
				if err != nil {
					toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
					fmt.Println(err)
					continue
				}

				questionMemo[item.Tid] = questionJudgeMemoType{
					judge:  questionJudge,
					detail: questionDetail,
				}
			}
			isContest, err := contest.IsContestSubmission(item.Sid)
			if err != nil {
				fmt.Println(err)
				toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid})
				continue
			}

			if err := submission.Starter(
				code,
				&item,
				questionMemo[item.Tid].judge,
				questionMemo[item.Tid].detail,
				isContest,
			); err != nil {
				fmt.Println(err)
				toBeRejected = append(toBeRejected, &toBeRejectObject{Sid: item.Sid, IsContest: isContest})
				continue
			}
		}

		fmt.Printf("[Judge Check] Total: %d, Rejected: %d \n",
			len(timeoutSubmissions),
			len(toBeRejected))
		if len(toBeRejected) > 0 {
			batchRejectSubmission(toBeRejected)
		}
	} else {
		//if _, err := db.DB.Table("submission").
		//	Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
		//	Update(&models.Submission{
		//		Status: "NO",
		//		Judge:  []models.JudgeResult{},
		//	}); err != nil {
		//
		//	fmt.Println(err)
		//}

		sidList := make([]*toBeRejectObject, len(timeoutSubmissions))
		for i, item := range timeoutSubmissions {
			sidList[i] = &toBeRejectObject{
				IsContest: false,
				Sid:       item.Sid,
			}
		}
		batchRejectSubmission(sidList)
	}
}

func batchRejectSubmission(sidList []*toBeRejectObject) {
	normalSidList, contestSidList, potentialSidList := make([]uint32, len(sidList)), make([]uint32, 0), make([]uint32, 0)
	for i, item := range sidList {
		normalSidList[i] = item.Sid
		if item.IsContest {
			contestSidList = append(contestSidList, item.Sid)
		} else {
			potentialSidList = append(potentialSidList, item.Sid)
		}
	}

	extraContestSidList, err := contest.BatchIsContestSubmission(potentialSidList)
	if err != nil {
		fmt.Println(err)
	} else {
		contestSidList = append(contestSidList, extraContestSidList...)
	}

	if _, err := db.DB.Table("submission").
		In("sid", sidList).
		Cols("status", "judge").
		Update(
			&models.Submission{
				Status: "NO",
				Judge:  []models.JudgeResult{},
			}); err != nil {

		fmt.Println(err)
	}

	if contestSidList != nil && len(contestSidList) > 0 {
		if _, err := db.DB.Table("contest_submission").
			In("sid", sidList).
			Update(
				&models.ContestSubmission{
					Status: -1,
				}); err != nil {

			fmt.Println(err)
		}
	}
}
