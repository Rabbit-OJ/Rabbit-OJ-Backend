package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"xorm.io/xorm"
)

func Edit(tid, subject, content string, difficulty uint8, timeLimit, spaceLimit uint32) error {
	_, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		questionOverview := models.Question{
			Subject:    subject,
			Difficulty: difficulty,
			TimeLimit:  timeLimit,
			SpaceLimit: spaceLimit,
		}

		if _, err := session.Table("question").
			Where("tid = ?", tid).
			Update(&questionOverview); err != nil {
			return nil, err
		}

		questionContent := models.QuestionContent{
			Content: content,
		}

		if _, err := session.Table("question_content").
			Where("tid = ?", tid).
			Update(&questionContent); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return err
}
