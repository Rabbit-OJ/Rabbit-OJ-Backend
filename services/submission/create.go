package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/question"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

func Create(tx *gorm.DB, tid, uid, language, fileName string) (*models.Submission, error) {
	questionJudge, err := question.JudgeInfo(tid)

	if err != nil {
		return nil, err
	}

	judgeArr := make([]models.JudgeResult, questionJudge.DatasetCount)
	for i := range judgeArr {
		judgeArr[i].Status = "ING"
	}
	judgeJSON, err := json.Marshal(judgeArr)

	if err != nil {
		return nil, err
	}

	submission := models.Submission{
		Tid:      tid,
		Uid:      uid,
		Language: language,
		FileName: fileName,
		Judge:    judgeJSON,
		Status:   "ING",
	}

	if err := tx.Create(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}
