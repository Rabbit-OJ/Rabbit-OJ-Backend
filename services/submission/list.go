package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func List(uid string, page int) ([]models.SubmissionLite, error) {
	var list []models.SubmissionLite

	err := db.DB.
		Select("`submission`.*, `question`.`subject` AS question_title").
		Table("submission").
		Join("INNER", "question", "`submission`.`tid` = `question`.`tid`").
		Where("`submission`.`uid` = ?", uid).
		Desc("`submission`.`sid`").
		Limit(config.PageSize, (page-1)*config.PageSize).Find(&list)

	if err != nil {
		return nil, err
	}
	return list, nil
}

func ListCount(uid string) (int64, error) {
	return db.DB.Table("submission").Where("uid = ?", uid).Count()
}
