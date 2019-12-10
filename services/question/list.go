package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(page uint32) ([]models.Question, error) {
	var list []models.Question

	err := db.DB.Table("question").
		Order("tid asc").
		Limit(config.PageSize).
		Offset((page - 1) * config.PageSize).
		Scan(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListCount() (uint32, error) {
	count := uint32(0)
	if err := db.DB.Table("question").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}