package submission

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/question"
	"time"
)

func Create(tid, uid uint32, language, fileName string) (*models.Submission, error) {
	questionJudge, err := question.JudgeInfo(tid)
	if err != nil {
		return nil, err
	}

	judgeArr := make([]models.JudgeResult, questionJudge.DatasetCount)
	for i := range judgeArr {
		judgeArr[i].Status = "ING"
	}

	submission := models.Submission{
		Tid:       tid,
		Uid:       uid,
		Language:  language,
		FileName:  fileName,
		Judge:     judgeArr,
		Status:    "ING",
		TimeUsed:  0,
		SpaceUsed: 0,
		CreatedAt: time.Now(),
	}

	_, err = db.DB.Table("submission").Insert(&submission)
	if err != nil {
		return nil, err
	}

	return &submission, nil
}
