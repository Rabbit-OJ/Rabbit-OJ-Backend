package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

const (
	StatusAC      = 1
	StatusPending = 0
	StatusERR     = -1
)

func Submit(sid, cid, uid, tid string, totalTime uint32) error {
	contestSubmission := models.ContestSubmission{
		Sid:       sid,
		Cid:       cid,
		Uid:       uid,
		Tid:       tid,
		Status:    StatusPending,
		TotalTime: totalTime,
	}

	return db.DB.Create(&contestSubmission).Error
}

func ChangeSubmitState(sid string, status int) error {
	return db.DB.
		Table("contest_submission").
		Where("sid = ?", sid).
		Update("status", status).Error
}
