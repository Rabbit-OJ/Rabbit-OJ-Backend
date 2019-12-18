package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(page uint32, showHide bool) ([]models.Question, error) {
	var list []models.Question
	var err error

	if showHide {
		err = db.DB.Table("question").
			Order("tid asc").
			Limit(config.PageSize).
			Offset((page - 1) * config.PageSize).
			Scan(&list).Error
	} else {
		err = db.DB.Table("question").
			Where("hide = ?", 0).
			Order("tid asc").
			Limit(config.PageSize).
			Offset((page - 1) * config.PageSize).
			Scan(&list).Error
	}

	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListCount(showHide bool) (uint32, error) {
	count := uint32(0)
	var err error

	if showHide {
		err = db.DB.Table("question").Count(&count).Error
	} else {
		err = db.DB.Table("question").
			Where("hide = ?", 0).
			Count(&count).Error
	}

	if err != nil {
		return 0, err
	}
	return count, nil
}
