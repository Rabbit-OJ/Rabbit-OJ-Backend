package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func JudgeInfo(tid string) (*models.QuestionJudge, error) {
	judge := &models.QuestionJudge{}

	if err := db.DB.Where("tid = ?", tid).First(&judge).Error; err != nil {
		return nil, err
	}

	return judge, nil
}
