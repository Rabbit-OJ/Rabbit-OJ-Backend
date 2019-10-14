package submission

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/utils"
)

func List(uid string, page uint32) ([]models.SubmissionLite, error) {
	var list []models.SubmissionLite

	err := db.DB.Table("submission").
		Where("uid = ?", uid).
		Limit(utils.PageSize).
		Offset((page - 1) * utils.PageSize).
		Scan(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}
