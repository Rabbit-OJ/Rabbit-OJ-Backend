package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
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

	return &contest, nil
}
