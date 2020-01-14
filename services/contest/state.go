package contest

import (
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
	DoingMap sync.Map
)

func CheckContestState(cid uint32) (int, error) {
	contest, err := Info(cid)
	if err != nil {
		return -1, err
	}

	now, start, end := time.Now().Unix(), time.Time(contest.StartTime).Unix(), time.Time(contest.EndTime).Unix()
	if now < start {
		return NotBegin, nil
	} else if start <= now && now <= end {
		return Running, nil
	} else if now > end {
		return Finished, nil
	}
	return -1, errors.New("contest arguments error")
}

func HavePendingSubmission(cid uint32) (bool, error) {
	count, err := db.DB.Table("contest_submission").
		Select("1").
		Where("cid = ? AND status = ?", cid, 0).
		Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func Finish(cid uint32) error {
	if _, ok := DoingMap.Load(cid); ok {
		return nil
	}

	DoingMap.Store(cid, 1)
	defer func() {
		DoingMap.Delete(cid)
	}()

	contestHub.RemoveContestHubAllContest(cid)
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

	if err := UpdateContestStatus(cid, RoundEnd); err != nil {
		return err
	}

	fmt.Printf("[Contest] #%d Ended \n", cid)
	return nil
}
