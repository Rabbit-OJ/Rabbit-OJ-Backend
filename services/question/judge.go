package question

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func JudgeInfo(tid uint32) (*models.QuestionJudge, error) {
	judge := models.QuestionJudge{}
	found, err := db.DB.Table("question_judge").Where("tid = ?", tid).Get(&judge)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("judge info not found")
	}

	return &judge, nil
}

func JudgeSet(tid uint32, form *forms.QuestionJudgeSetForm) error {
	questionJudge := models.QuestionJudge{
		Mode:         form.Mode,
		DatasetCount: form.DatasetCount,
		Version:      form.Version,
	}

	if _, err := db.DB.Table("question_judge").
		Where("tid = ?", tid).
		Cols("mode", "dataset_count", "version").
		Update(&questionJudge); err != nil {
		return err
	}

	return nil
}
