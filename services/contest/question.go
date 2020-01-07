package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func Question(cid uint32) ([]models.ContestQuestion, error) {
	var contestQuestion []models.ContestQuestion

	if err := db.DB.
		Table("contest_question").
		Where("cid = ?", cid).
		Asc("id").Find(&contestQuestion);
		err != nil {
		return nil, err
	}

	return contestQuestion, nil
}

func QuestionOne(cid, tid uint32) (*models.ContestQuestion, error) {
	var contestQuestion models.ContestQuestion

	found, err := db.DB.
		Table("contest_question").
		Where("cid = ? AND tid = ?", cid, tid).
		Get(&contestQuestion)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("question doesn't exist")
	}

	return &contestQuestion, nil
}

func QuestionExtended(cid uint32) ([]models.ContestQuestionExtended, error) {
	var contestQuestionExtended []models.ContestQuestionExtended

	if err := db.DB.
		Table("contest_question").
		Select("contest_question.*, question.*, question_content.*").
		Join("INNER", "question", "`contest_question`.`tid` = `question`.`tid`").
		Join("INNER", "question_content", "`contest_question`.`tid` = `question_content`.`tid`").
		Where("cid = ?", cid).
		Asc("id").Find(&contestQuestionExtended); err != nil {
		return nil, err
	}

	//for i := range contestQuestionExtended {
	//	samplePtr := &contestQuestionExtended[i].SampleJSON
	//	if err := json.Unmarshal(contestQuestionExtended[i].Sample, samplePtr); err != nil {
	//		return nil, err
	//	}
	//}

	return contestQuestionExtended, nil
}

func QuestionMapTidToId(cid uint32) (map[uint32]int, []models.ContestQuestion, error) {
	var contestQuestion []models.ContestQuestion

	if err := db.DB.
		Table("contest_question").
		Where("cid = ?", cid).
		Asc("id").
		Find(&contestQuestion); err != nil {
		return nil, nil, err
	}

	questionMap := make(map[uint32]int)
	for _, item := range contestQuestion {
		questionMap[item.Tid] = item.Id
	}
	return questionMap, contestQuestion, nil
}
