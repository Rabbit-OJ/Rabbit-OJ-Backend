package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

const (
	RoundWaiting  = 0
	RoundStarting = 1
	RoundEnd      = 2
)

func Info(cid string) (*models.Contest, error) {
	contest := models.Contest{}
	if err := db.DB.Table("contest").
		Where("cid = ?", cid).
		First(&contest).Error; err != nil {
		return nil, err
	}

	return &contest, nil
}
