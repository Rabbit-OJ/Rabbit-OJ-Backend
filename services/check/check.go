package check

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

func checkRoutine(interval int, ctx context.Context) {
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Second):
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
			Find(&timeoutSubmissions).Error; err != nil {
			fmt.Println(err)
			return
		}

		toBeRejected, questionMemo := make([]string, 0), make(map[string]questionJudgeMemoType)
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
		if err := db.DB.Table("submission").
			Where("status = ? AND created_at <= ?", "ING", someMinutesBefore).
			Updates(map[string]string{"status": "NO", "judge": "[]"}).
			Error; err != nil {

			fmt.Println(err)
		}
	}
}

func batchRejectSubmission(sidList []string) {
	if err := db.DB.Table("submission").
		Where("sid in (?)", sidList).
		Updates(map[string]string{"status": "NO", "judge": "[]"}).
		Error; err != nil {

		fmt.Println(err)
	}
}
