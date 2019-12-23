package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
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
