package question

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

const InvalidTid = ""

func Create(uid, subject, content string, difficulty uint8, timeLimit, spaceLimit uint32) (string, error) {
	tx := db.DB.Begin()

	questionOverview := &models.Question{
		Uid:        uid,
		Subject:    subject,
		Difficulty: difficulty,
		TimeLimit:  timeLimit,
		SpaceLimit: spaceLimit,
	}

	if err := tx.Create(&questionOverview).Error; err != nil {
		tx.Rollback()
		return InvalidTid, err
	}

	questionContent := &models.QuestionContent{
		Tid:     questionOverview.Tid,
		Content: content,
	}

	if err := tx.Create(&questionContent).Error; err != nil {
		tx.Rollback()
		return InvalidTid, err
	}

	tx.Commit()
	return questionOverview.Tid, nil
}
