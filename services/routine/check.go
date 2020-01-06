package routine

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
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

	if config.Global.Extensions.CheckJudge.Requeue {
		var timeoutSubmissions []models.Submission
		if err := db.DB.Table("submission").
			Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
			Find(&timeoutSubmissions); err != nil {
			fmt.Println(err)
			return
		}

		toBeRejected, questionMemo := make([]uint32, 0), make(map[uint32]questionJudgeMemoType)
		for _, item := range timeoutSubmissions {
			path, err := files.CodePath(item.FileName)
			if err != nil {
				toBeRejected = append(toBeRejected, item.Sid)
				continue
			}
			exist := files.Exists(path)
			if !exist {
				toBeRejected = append(toBeRejected, item.Sid)
				continue
			}
			code, err := ioutil.ReadFile(path)
			if err != nil {
				toBeRejected = append(toBeRejected, item.Sid)
				continue
			}
			if _, ok := questionMemo[item.Tid]; !ok {
				questionDetail, err := question.Detail(item.Tid)
				if err != nil {
					toBeRejected = append(toBeRejected, item.Sid)
					continue
				}
				questionJudge, err := question.JudgeInfo(item.Tid)
				if err != nil {
					toBeRejected = append(toBeRejected, item.Sid)
					continue
				}

				questionMemo[item.Tid] = questionJudgeMemoType{
					judge:  questionJudge,
					detail: questionDetail,
				}
			}

			if err := submission.Starter(
				code,
				&item,
				questionMemo[item.Tid].judge,
				questionMemo[item.Tid].detail,
				false, // todo: test if it is a contest submission
			); err != nil {
				fmt.Println(err)
				toBeRejected = append(toBeRejected, item.Sid)
			}
		}

		fmt.Printf("[Judge Check] Total: %d, Rejected: %d \n",
			len(timeoutSubmissions),
			len(toBeRejected))

		if len(toBeRejected) > 0 {
			batchRejectSubmission(toBeRejected)
		}
	} else {
		if _, err := db.DB.Table("submission").
			Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
			Update(&models.Submission{
				Status: "NO",
				Judge:  []byte("[]"),
			}); err != nil {

			fmt.Println(err)
		}
	}
}

func batchRejectSubmission(sidList []uint32) {
	if _, err := db.DB.Table("submission").
		Where("sid in (?)", sidList).
		Update(
			&models.Submission{
				Status: "NO",
				Judge:  []byte("[]"),
			}); err != nil {

		fmt.Println(err)
	}
}
