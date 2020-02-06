package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"time"
	"xorm.io/xorm"
)

func Edit(cid string, form *forms.ContestEditForm) error {
	contestInfo := models.Contest{
		Name:      form.Name,
		Status:    form.Status,
		Penalty:   form.Penalty,
		StartTime: models.JSONTime(time.Unix(form.StartTime, 0)),
		EndTime:   models.JSONTime(time.Unix(form.EndTime, 0)),
		BlockTime: models.JSONTime(time.Unix(form.BlockTime, 0)),
	}

	if _, err := db.DB.Table("contest").
		Where("cid = ?", cid).
		Cols("name", "status", "penalty", "start_time", "end_time", "block_time").
		Update(&contestInfo); err != nil {
		return err
	}

	return nil
}

func EditQuestions(cid uint32, newQuestionList []forms.ContestQuestionEditItem) error {
	_, err := db.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		current, err := Question(cid)
		if err != nil {
			return nil, err
		}

		currentLen, newLen, recyclePtr := len(current), len(newQuestionList), 0
		for ; recyclePtr < newLen; recyclePtr++ {
			if recyclePtr < currentLen {
				if _, err := db.DB.Table("contest_question").
					Where("cqid = ?", current[recyclePtr].Cqid).
					Cols("tid", "id", "score").
					Update(&newQuestionList[recyclePtr]); err != nil {
					return nil, err
				}
			} else {
				if _, err := db.DB.Table("contest_question").
					Cols("tid", "id", "score").
					Insert(&newQuestionList[recyclePtr]); err != nil {
					return nil, err
				}
			}
		}

		toBeDeleted := make([]uint32, 0)
		for ; recyclePtr < currentLen; recyclePtr++ {
			toBeDeleted = append(toBeDeleted, current[recyclePtr].Cqid)
		}
		if len(toBeDeleted) > 0 {
			if _, err := db.DB.Table("contest_question").
				In("cqid", toBeDeleted).
				Delete(models.ContestQuestion{}); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
