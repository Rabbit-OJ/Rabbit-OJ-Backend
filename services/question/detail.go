package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

func Detail(tid string) (*models.QuestionDetail, error) {
	content := models.QuestionDetail{}

	if err := db.DB.Table("question_content").
		Select("`question`.`tid`, content, subject, hide, attempt, accept, difficulty, time_limit, sample, space_limit, created_at").
		Joins("INNER JOIN question ON question.tid = question_content.tid").
		Where("`question`.`tid` = ?", tid).
		First(&content).
		Error; err != nil {
		return nil, err
	} else {
		return &content, nil
	}
}
