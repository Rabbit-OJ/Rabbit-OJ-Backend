package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"xorm.io/xorm"
)

func Edit(tid string, form *forms.QuestionForm) error {
	_, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		questionOverview := models.Question{
			Subject:    form.Subject,
			Difficulty: form.Difficulty,
			TimeLimit:  form.TimeLimit,
			SpaceLimit: form.SpaceLimit,
			Hide:       form.Hide,
		}
		if _, err := session.Table("question").
			Where("tid = ?", tid).
			Cols("subject", "difficulty", "time_limit", "space_limit", "hide").
			Update(&questionOverview); err != nil {
			return nil, err
		}

		questionContent := models.QuestionContent{
			Content: form.Content,
			Sample:  form.Sample,
		}

		if _, err := session.Table("question_content").
			Where("tid = ?", tid).
			Cols("content", "sample").
			Update(&questionContent); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}
