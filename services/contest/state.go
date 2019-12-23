package contest

import (
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	NotBegin = 0
	Running  = 1
	Finished = 2
)

var (
	FinishLock sync.Mutex
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

func Start(cid string) error {
	if err := ChangeContestState(cid, Running); err != nil {
		return err
	}

	return nil
}

func HavePendingSubmission(cid string) (bool, error) {
	count := 0

	if err := db.DB.Table("contest_submission").
		Select("1").
		Where("cid = ? AND status = ?", cid, 0).
		Count(&count).
		Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func Finish(cid string) error {
	FinishLock.Lock()
	defer FinishLock.Unlock()

	websocket.SocketHub.ContestHub.RemoveContestHubAllContest(cid)
	for {
		have, err := HavePendingSubmission(cid)

		if err != nil {
			fmt.Println(err)
			return err
		}
		if !have {
			break
		}

		time.Sleep(15 * time.Second)
	}

	if err := ChangeContestState(cid, Finished); err != nil {
		return err
	}
	return nil
}
