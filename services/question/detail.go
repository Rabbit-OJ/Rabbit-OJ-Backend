package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func Detail(tid uint32) (*models.QuestionDetail, error) {
	content := models.QuestionDetail{}

	found, err := db.DB.Table("question_content").
		Join("INNER", "question", "question.tid = question_content.tid").
		Select("`question`.`tid`, content, subject, hide, attempt, accept, difficulty, time_limit, sample, space_limit, created_at").
		Where("`question`.`tid` = ?", tid).
		Get(&content)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("question doesn't exist")
	}

	return &content, nil
}
