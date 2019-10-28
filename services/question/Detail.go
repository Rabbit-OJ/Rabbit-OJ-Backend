package question

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Detail(tid string) (*models.QuestionDetail, error) {
	content := &models.QuestionDetail{}

	if err := db.DB.Table("question_content").
		Select("`question`.`tid`, content, subject, attempt, accept, difficulty, time_limit, space_limit, created_at").
		Joins("INNER JOIN question ON question.tid = question_content.tid").
		Where("`question`.`tid` = ?", tid).
		First(content).
		Error; err != nil {
		return nil, err
	} else {
		return content, nil
	}
}
