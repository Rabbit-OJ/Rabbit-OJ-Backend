package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func Edit(tid, subject, content string, difficulty uint8, timeLimit, spaceLimit uint32) error {
	tx := db.DB.Begin()

	questionOverview := &models.Question{
		Subject:    subject,
		Difficulty: difficulty,
		TimeLimit:  timeLimit,
		SpaceLimit: spaceLimit,
	}

	if err := tx.Where("tid = ?", tid).Update(&questionOverview).Error; err != nil {
		tx.Rollback()
		return err
	}

	questionContent := &models.QuestionContent{
		Content: content,
	}

	if err := tx.Where("tid = ?", tid).Update(&questionContent).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
