package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func Update(sid, timeUsed, spaceUsed uint32, status string, judge []models.JudgeResult) error {
	updateObj := models.Submission{
		TimeUsed:  timeUsed,
		SpaceUsed: spaceUsed,
		Status:    status,
		Judge:     judge,
	}

	if _, err := db.DB.Table("submission").
		Where("sid = ?", sid).
		Update(&updateObj); err != nil {
		return err
	}

	return nil
}
