package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"xorm.io/xorm"
)

func RegenerateUserScore(session *xorm.Session, cid, uid uint32, isAccepted bool) error {
	contestInfo, err := Info(cid)
	if err != nil {
		return err
	}

	questionMapTidToId, contestQuestion, err := QuestionMapTidToId(cid)
	if err != nil {
		return err
	}

	progress := make([]responses.ScoreBoardProgress, len(questionMapTidToId))
	var contestSubmissionList []models.ContestSubmission

	if err := session.Table("contest_submission").
		Where("cid = ? AND uid = ?", cid, uid).
		Find(&contestSubmissionList); err != nil {
		return err
	}

	// calc Number
	for _, item := range contestSubmissionList {
		questionId := questionMapTidToId[item.Tid]

		if item.Status == StatusPending {
			continue
		} else if item.Status == StatusAC {
			progress[questionId].Status = StatusAC

			totalTime := &progress[questionId].TotalTime
			if *totalTime == 0 || *totalTime > item.TotalTime {
				*totalTime = item.TotalTime
			}
		} else if item.Status == StatusERR {
			progress[questionId].Bug++
		}
	}

	score, totalTime := uint32(0), uint32(0)
	// calc penalty
	for i, item := range progress {
		if item.Status == StatusAC {
			score += contestQuestion[i].Score

			currentTime := item.TotalTime + contestInfo.Penalty*item.Bug
			if currentTime > totalTime {
				totalTime = currentTime
			}
		}
	}

	if isAccepted {
		if _, err := db.DB.Table("contest_user").
			Where("cid = ? AND uid = ?", cid, uid).
			Update(&models.ContestUser{
				Score:     score,
				TotalTime: totalTime,
			}); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
