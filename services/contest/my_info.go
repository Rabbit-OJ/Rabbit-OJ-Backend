package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/db"
)

func User(uid, cid string) (*models.ContestUser, error) {
	contestUser := models.ContestUser{}

	if err := db.DB.Table("contest_user").
		Where("uid = ? AND cid = ?", uid, cid).
		First(&contestUser).Error; err != nil {
		return nil, err
	}

	return &contestUser, nil
}

func QueryUserRank(uid, cid string) (uint32, error) {
	var rank struct {
		Rank uint32 `gorm:"rank"`
	}

	if err := db.DB.
		Raw("SELECT `rank` FROM (SELECT uid, RANK() OVER "+
			"(ORDER BY score DESC, total_time ASC) "+
			"`rank` FROM contest_user WHERE cid = ?) AS temp "+
			"WHERE uid = ?", cid, uid).
		Scan(&rank).Error; err != nil {
		return 0, err
	}

	return rank.Rank, nil
}

func MyInfo(uid, cid string, contest *models.Contest) (*responses.ContestMyInfo, error) {
	contestMyInfo := responses.ContestMyInfo{
		Rank:      0,
		TotalTime: 0,
		Score:     0,
		Progress:  make([]responses.ScoreBoardProgress, contest.Count),
	}

	contestUser, err := User(uid, cid)
	if err != nil {
		return nil, err
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
		Find(&contestSubmissionList).
		Error; err != nil {
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

	// calc penalty
	for j := range contestMyInfo.Progress {
		if contestMyInfo.Progress[j].Status == StatusAC {
			contestMyInfo.Progress[j].TotalTime += int64(contestMyInfo.Progress[j].Bug) * contest.Penalty
		}
	}

	return &contestMyInfo, nil
}
