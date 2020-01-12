package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/db"
	"errors"
	"strconv"
)

func User(uid, cid uint32) (*models.ContestUser, bool, error) {
	contestUser := models.ContestUser{}

	found, err := db.DB.Table("contest_user").
		Where("uid = ? AND cid = ?", uid, cid).
		Get(&contestUser)

	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	return &contestUser, true, nil
}

func QueryUserRank(uid, cid uint32) (uint32, error) {
	results, err := db.DB.
		QueryString("SELECT `rank` FROM (SELECT uid, RANK() OVER "+
			"(ORDER BY score DESC, total_time ASC) "+
			"`rank` FROM contest_user WHERE cid = ?) AS temp "+
			"WHERE uid = ?", cid, uid)

	if err != nil {
		return 0, err
	}
	if len(results) <= 0 {
		return 0, errors.New("user doesn't exist")
	}

	rankNumber, err := strconv.ParseUint(results[0]["rank"], 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(rankNumber), nil
}

func MyInfo(uid, cid uint32, contest *models.Contest) (*responses.ContestMyInfo, error) {
	contestMyInfo := responses.ContestMyInfo{
		Rank:       0,
		TotalTime:  0,
		Score:      0,
		Progress:   make([]responses.ScoreBoardProgress, contest.Count),
		Registered: true,
	}

	contestUser, found, err := User(uid, cid)
	if err != nil {
		return nil, err
	}
	if !found {
		contestMyInfo.Registered = false
		return &contestMyInfo, nil
	}

	contestMyInfo.TotalTime, contestMyInfo.Score = contestUser.TotalTime, contestUser.Score
	inBlockTime, err := InBlockTime(cid)
	if err != nil {
		return nil, err
	}

	if !inBlockTime {
		rank, err := QueryUserRank(uid, cid)
		if err != nil {
			return nil, err
		}
		contestMyInfo.Rank = rank
	}

	questionMapTidToId, _, err := QuestionMapTidToId(cid)
	if err != nil {
		return nil, err
	}

	var contestSubmissionList []models.ContestSubmission
	if err := db.DB.Table("contest_submission").
		Where("cid = ? AND uid = ?", cid, uid).
		Find(&contestSubmissionList); err != nil {
		return nil, err
	}

	// calc Number
	for _, item := range contestSubmissionList {
		questionId := questionMapTidToId[item.Tid]

		if item.Status == StatusPending {
			continue
		} else if item.Status == StatusAC {
			contestMyInfo.Progress[questionId].Status = StatusAC

			totalTime := &contestMyInfo.Progress[questionId].TotalTime
			if *totalTime == 0 || *totalTime > item.TotalTime {
				*totalTime = item.TotalTime
			}
		} else if item.Status == StatusERR {
			contestMyInfo.Progress[questionId].Bug++
		}
	}

	//// calc penalty
	//for j := range contestMyInfo.Progress {
	//	if contestMyInfo.Progress[j].Status == StatusAC {
	//		contestMyInfo.Progress[j].TotalTime += contestMyInfo.Progress[j].Bug * contest.Penalty
	//	}
	//}

	return &contestMyInfo, nil
}
