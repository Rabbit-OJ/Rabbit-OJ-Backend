package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(page uint32) ([]models.Contest, error) {
	var list []models.Contest

	if err := db.DB.Table("contest").
		Order("cid asc").
		Limit(config.PageSize).
		Offset((page - 1) * config.PageSize).
		Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func ListCount() (uint32, error) {
	count := uint32(0)

	if err := db.DB.Table("contest").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
