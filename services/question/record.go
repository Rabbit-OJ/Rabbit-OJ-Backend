package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func Record(uid, tid string, page uint32) ([]models.SubmissionLite, error) {
	var list []models.SubmissionLite

	err := db.DB.
		Select("`submission`.*, `question`.`subject` AS question_title").
		Table("submission").
		Joins("INNER JOIN question ON `submission`.`tid` = `question`.`tid`").
		Where("`submission`.`uid` = ? AND `submission`.`tid` = ?", uid, tid).
		Order("`submission`.`sid` DESC").
		Limit(config.PageSize).
		Offset((page - 1) * config.PageSize).
		Scan(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func RecordCount(uid, tid string) (uint32, error) {
	count := uint32(0)
	if err := db.DB.Table("submission").
		Where("uid = ? AND tid = ?", uid, tid).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
