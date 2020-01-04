package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"encoding/json"
)

func Question(cid string) ([]models.ContestQuestion, error) {
	var contestQuestion []models.ContestQuestion

	if err := db.DB.
		Table("contest_question").
		Where("cid = ?", cid).
		Order("id asc").
		Scan(&contestQuestion).Error; err != nil {
		return nil, err
	}

	return contestQuestion, nil
}

func QuestionOne(cid, id string) (*models.ContestQuestion, error) {
	var contestQuestion *models.ContestQuestion

	if err := db.DB.
		Table("contest_question").
		Where("cid = ? AND id = ?", cid, id).
		First(contestQuestion).Error; err != nil {
		return nil, err
	}

	return contestQuestion, nil
}

func QuestionExtended(cid string) ([]models.ContestQuestionExtended, error) {
	var contestQuestionExtended []models.ContestQuestionExtended

	if err := db.DB.
		Table("contest_question").
		Select("contest_question.*, question.*, question_content.*").
		Joins("INNER JOIN question ON `contest_question`.`tid` = `question`.`tid`").
		Joins("INNER JOIN question_content ON `contest_question`.`tid` = `question_content`.`tid`").
		Where("cid = ?", cid).
		Order("id asc").
		Scan(&contestQuestionExtended).Error; err != nil {
		return nil, err
	}

	for i := range contestQuestionExtended {
		samplePtr := &contestQuestionExtended[i].SampleJSON
		if err := json.Unmarshal(contestQuestionExtended[i].Sample, samplePtr); err != nil {
			return nil, err
		}
	}

	return contestQuestionExtended, nil
}

func QuestionMapTidToId(cid string) (map[string]int, []models.ContestQuestion, error) {
	var contestQuestion []models.ContestQuestion

	if err := db.DB.
		Table("contest_question").
		Where("cid = ?", cid).
		Order("id asc").
		Scan(&contestQuestion).Error; err != nil {
		return nil, nil, err
	}

	questionMap := make(map[string]int)
	for _, item := range contestQuestion {
		questionMap[item.Tid] = item.Id
	}
	return questionMap, contestQuestion, nil
}
