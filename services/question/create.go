package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"time"
	"xorm.io/xorm"
)

const InvalidTid = 0

func Create(uid uint32, subject, content string, difficulty uint8, timeLimit, spaceLimit uint32) (int, error) {
	tid, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		questionOverview := models.Question{
			Uid:        uid,
			Subject:    subject,
			Difficulty: difficulty,
			TimeLimit:  timeLimit,
			SpaceLimit: spaceLimit,
			CreatedAt:  time.Now(),
		}

		if _, err := session.Table("question").Insert(&questionOverview); err != nil {
			return InvalidTid, err
		}

		questionContent := models.QuestionContent{
			Content: content,
		}

		if _, err := session.Table("question_content").Insert(&questionContent); err != nil {
			return InvalidTid, err
		}

		return questionOverview.Tid, nil
	})

	if val, ok := tid.(int); ok {
		return val, err
	} else {
		return InvalidTid, err
	}
}
