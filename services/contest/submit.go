package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"time"
	"xorm.io/xorm"
)

const (
	StatusAC      = 1
	StatusPending = 0
	StatusERR     = -1
)

func Submit(sid, cid, uid, tid, totalTime uint32) error {
	contestSubmission := models.ContestSubmission{
		Sid:       sid,
		Cid:       cid,
		Uid:       uid,
		Tid:       tid,
		Status:    StatusPending,
		TotalTime: totalTime,
		CreatedAt: time.Now(),
	}

	_, err := db.DB.Table("contest_submission").Insert(&contestSubmission)
	return err
}

func ChangeSubmitState(session *xorm.Session, sid uint32, status int) error {
	_, err := session.
		Table("contest_submission").
		Where("sid = ?", sid).
		Cols("status").
		Update(&models.ContestSubmission{
			Status: status,
		})

	return err
}
