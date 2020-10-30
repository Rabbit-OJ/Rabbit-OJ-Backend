package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
)

func Record(uid, tid uint32, page int) ([]models.SubmissionLite, error) {
	list := make([]models.SubmissionLite, 0)

	err := db.DB.
		Select("`submission`.*, `question`.`subject` AS question_title").
		Table("submission").
		Join("INNER", "question", "`submission`.`tid` = `question`.`tid`").
		Desc("`submission`.`sid`").
		Where("`submission`.`uid` = ? AND `submission`.`tid` = ?", uid, tid).
		Limit(config.PageSize, (page-1)*config.PageSize).
		Find(&list)

	if err != nil {
		return nil, err
	}
	return list, nil
}

func RecordCount(uid, tid uint32) (int64, error) {
	count, err := db.DB.Table("submission").
		Where("uid = ? AND tid = ?", uid, tid).
		Count()

	if err != nil {
		return 0, err
	}
	return count, nil
}
