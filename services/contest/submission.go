package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"xorm.io/xorm"
)

func SubmissionList(uid, cid uint32) ([]models.ContestSubmission, error) {
	var contestSubmissionList []models.ContestSubmission
	if err := db.DB.
		Where("`contest_submission`.cid = ? AND `contest_submission`.uid = ?", cid, uid).
		Table("contest_submission").
		Join("INNER", "submission", "`contest_submission`.`sid` = `submission`.`sid`").
		Desc("`contest_submission`.sid").Find(&contestSubmissionList);
		err != nil {
		return nil, err
	}

	return contestSubmissionList, nil
}

func SubmissionInfo(session *xorm.Session, sid uint32) (*models.ContestSubmission, error) {
	var contestSubmission models.ContestSubmission

	found, err := session.
		Table("contest_submission").
		Where("sid = ?", sid).
		Get(&contestSubmission)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("submission doesn't exist")
	}

	return &contestSubmission, nil
}
