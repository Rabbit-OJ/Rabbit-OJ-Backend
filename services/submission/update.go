package submission

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Update(sid string, timeUsed, spaceUsed uint32, status string, judge []byte) error {
	updateObj := &models.Submission{
		TimeUsed:  timeUsed,
		SpaceUsed: spaceUsed,
		Status:    status,
		Judge:     judge,
	}

	if err := db.DB.Table("submission").Where("sid = ?", sid).Update(updateObj).Error; err != nil {
		return err
	}

	return nil
}
