package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"github.com/jinzhu/gorm"
)

const (
	StatusAC      = 1
	StatusPending = 0
	StatusERR     = -1
)

func Submit(sid, cid, uid, tid string, totalTime int64) error {
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

func ChangeSubmitState(tx *gorm.DB, sid string, status int) error {
	return tx.
		Table("contest_submission").
		Where("sid = ?", sid).
		Update("status", status).Error
}
