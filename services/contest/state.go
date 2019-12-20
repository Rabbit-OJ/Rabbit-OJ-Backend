package contest

import (
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"time"
)

const (
	NotBegin = 0
	Running  = 1
	Finished = 2
)

func ChangeContestState(cid string, status int) error {
	return db.DB.
		Table("contest").
		Where("cid = ?", cid).
		Update("status", status).Error
}

func CheckContestState(cid string) (int, error) {
	contest, err := Info(cid)
	if err != nil {
		return -1, err
	}

	now, start, end := time.Now().Unix(), contest.StartTime.Unix(), contest.EndTime.Unix()
	if now < start {
		return NotBegin, nil
	} else if start <= now && now <= end {
		return Running, nil
	} else if now > end {
		return Finished, nil
	}
	return -1, errors.New("contest arguments error")
}
