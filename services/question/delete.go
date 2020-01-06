package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"xorm.io/xorm"
)

// NOTE: DELETE a problem WILL NOT CASCADE DELETE its submission records and codes, do it manually!

func Delete(tid uint32) error {
	_, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		if _, err := session.Table("question").
			Delete(&models.Question{
				Tid: tid,
			}); err != nil {
			return nil, err
		}

		if _, err := session.Table("question_content").
			Delete(&models.QuestionContent{
				Tid: tid,
			}); err != nil {
			return nil, err
		}

		if _, err := session.Table("question_judge").
			Delete(&models.QuestionJudge{
				Tid: tid,
			}); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
