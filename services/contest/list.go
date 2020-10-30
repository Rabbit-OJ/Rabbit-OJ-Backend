package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(page int) ([]models.Contest, error) {
	list := make([]models.Contest, 0)

	if err := db.DB.Table("contest").
		Asc("cid").
		Limit(config.PageSize, (page-1)*config.PageSize).
		Find(&list); err != nil {

		return nil, err
	}

	return list, nil
}

func ListCount() (int64, error) {
	count, err := db.DB.Table("contest").Count()

	if err != nil {
		return 0, err
	}
	return count, nil
}
