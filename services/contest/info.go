package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"fmt"
	"time"
)

const (
	RoundWaiting     = 0
	RoundStarting    = 1
	RoundEnd         = 2
	RoundCalculating = 3
)

func Info(cid uint32) (*models.Contest, error) {
	contest := models.Contest{}
	found, err := db.DB.Table("contest").
		Where("cid = ?", cid).
		Get(&contest)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("contest doesn't exist")
	}

	if err := UpdateInfo(&contest); err != nil {
		return nil, err
	}
	return &contest, nil
}

func UpdateInfo(contest *models.Contest) error {
	if contest.Status == RoundCalculating || contest.Status == RoundEnd {
		return nil
	}

	previousStatus := contest.Status
	start, now, end := time.Time(contest.StartTime).Unix(), time.Now().Unix(), time.Time(contest.EndTime).Unix()
	if start <= now && now <= end {
		contest.Status = RoundStarting
	} else if now > end && contest.Status == RoundStarting {
		contest.Status = RoundCalculating
	}

	if previousStatus != contest.Status {
		_ = UpdateContestStatus(contest.Cid, contest.Status)

		if contest.Status == RoundWaiting {
			fmt.Printf("[Contest] #%d Starting \n", contest.Cid)
		} else if contest.Status == RoundCalculating {
			fmt.Printf("[Contest] #%d Calculating \n", contest.Cid)
			go func() {
				err := Finish(contest.Cid)
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
	}
	return nil
}

func UpdateContestStatus(cid uint32, status int32) error {
	_, err := db.DB.
		Table("contest").
		Where("cid = ?", cid).
		Cols("status").
		Update(&models.Contest{
			Status: status,
		})

	return err
}
