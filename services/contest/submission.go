package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"github.com/jinzhu/gorm"
)

func SubmissionList(uid, cid string) ([]models.ContestSubmission, error) {
	var contestSubmissionList []models.ContestSubmission
	if err := db.DB.
		Where("`contest_submission`.cid = ? AND `contest_submission`.uid = ?", cid, uid).
		Table("contest_submission").
		Joins("INNER JOIN submission ON `contest_submission`.`sid` = `submission`.`sid`").
		Order("`contest_submission`.sid DESC").
		Scan(&contestSubmissionList).
		Error; err != nil {
		return nil, err
	}

	return contestSubmissionList, nil
}

func SubmissionInfo(tx *gorm.DB, sid string) (*models.ContestSubmission, error) {
	var contestSubmission models.ContestSubmission
	if err := tx.
		Table("contest_submission").
		Where("sid = ?", sid).
		First(&contestSubmission).
		Error; err != nil {
		return nil, err
	}

	return &contestSubmission, nil
}