package submission

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Detail (sid string) (*models.Submission, error) {
	submission := &models.Submission{}

	if err := db.DB.
		Select("`submission`.*, `question`.`subject` AS question_title").
		Where("sid = ?", sid).
		Table("submission").
		Joins("INNER JOIN question ON `submission`.`tid` = `question`.`tid`").
		First(&submission).Error; err != nil {
		return nil, err
	} else {
		return submission, nil
	}
}