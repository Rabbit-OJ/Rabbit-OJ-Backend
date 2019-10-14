package submission

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Detail (sid string) (*models.Submission, error) {
	submission := &models.Submission{}

	if err := db.DB.Where("sid = ?", sid).First(&submission).Error; err != nil {
		return nil, err
	} else {
		return submission, nil
	}
}