package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func Content(tid uint32) (*models.QuestionContent, error) {
	content := models.QuestionContent{}

	found, err := db.DB.Where("tid = ?", tid).Get(&content)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("question content doesn't exist")
	}
	return &content, nil
}
