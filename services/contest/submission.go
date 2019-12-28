package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
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
