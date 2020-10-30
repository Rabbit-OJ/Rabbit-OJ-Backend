package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"strconv"
	"xorm.io/xorm"
)

func SubmissionList(uid, cid uint32) ([]models.ContestSubmission, error) {
	contestSubmissionList := make([]models.ContestSubmission, 0)
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

func SubmissionOne(sid uint32) (*models.ContestSubmission, error) {
	session := db.DB.NewSession()
	defer session.Close()

	return SubmissionInfo(session, sid)
}

func IsContestSubmission(sid uint32) (bool, error) {
	var contestSubmission models.ContestSubmission

	found, err := db.DB.
		Table("contest_submission").
		Where("sid = ?", sid).
		Get(&contestSubmission)

	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	return true, nil
}

func BatchIsContestSubmission(sidList []uint32) ([]uint32, error) {
	if resultMap, err := db.DB.Table("contest_submission").
		Select("sid").
		In("sid", sidList).
		QueryString(); err != nil {
		return nil, err
	} else {
		contestSidList := make([]uint32, len(resultMap))
		for i, item := range resultMap {
			sid, err := strconv.ParseUint(item["sid"], 10, 32)
			if err != nil {
				return nil, err
			}

			contestSidList[i] = uint32(sid)
		}

		return contestSidList, nil
	}
}
