package submission

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Detail (sid string) (*models.SubmissionExtended, error) {
	submission := &models.SubmissionExtended{}

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