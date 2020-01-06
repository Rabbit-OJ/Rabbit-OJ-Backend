package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(page int, showHide bool) ([]models.Question, error) {
	var list []models.Question
	var err error

	if showHide {
		err = db.DB.Table("question").
			Asc("tid").Limit(config.PageSize, (page-1)*config.PageSize).
			Find(&list)
	} else {
		err = db.DB.Table("question").
			Where("hide = ?", 0).
			Asc("tid").Limit(config.PageSize, (page-1)*config.PageSize).
			Find(&list)
	}

	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListCount(showHide bool) (int64, error) {
	var err error
	var count int64

	if showHide {
		count, err = db.DB.Table("question").Count()
	} else {
		count, err = db.DB.Table("question").
			Where("hide = ?", 0).
			Count()
	}

	if err != nil {
		return 0, err
	}
	return count, nil
}
