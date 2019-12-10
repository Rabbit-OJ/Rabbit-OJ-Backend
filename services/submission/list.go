package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(uid string, page uint32) ([]models.SubmissionLite, error) {
	var list []models.SubmissionLite

	err := db.DB.
		Select("`submission`.*, `question`.`subject` AS question_title").
		Table("submission").
		Joins("INNER JOIN question ON `submission`.`tid` = `question`.`tid`").
		Where("`submission`.`uid` = ?", uid).
		Order("`submission`.`sid` DESC").
		Limit(config.PageSize).
		Offset((page - 1) * config.PageSize).
		Scan(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListCount(uid string) (uint32, error) {
	count := uint32(0)
	if err := db.DB.Table("submission").Where("uid = ?", uid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}