package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
	"time"
)

func InBlockTime(cid string) (bool, error) {
	contest, err := Info(cid)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()
	return contest.BlockTime.Unix() <= now && now <= contest.EndTime.Unix(), nil
}

func ScoreBoard(cid string, page uint32) ([]responses.ScoreBoard, error) {
	var scoreBoard []responses.ScoreBoard

	contest, err := Info(cid)
	if err != nil {
		return nil, err
	}

	questionMapTidToId, _, err := QuestionMapTidToId(cid)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Raw("SELECT `user`.username, `user`.uid, `rank`, score, total_time "+
		"FROM (SELECT cid, uid, score, total_time, RANK() OVER (ORDER BY score DESC, total_time ASC) `rank` "+
		"FROM contest_user) AS temp "+
		"INNER JOIN `user` ON `temp`.`uid`=`user`.`uid` "+
		"WHERE cid = ? LIMIT ?, ?",
		cid, (page-1)*config.PageSize, page*config.PageSize).
		Scan(&scoreBoard).Error; err != nil {
		return nil, err
	}

	uidList, mapUidToIndex := make([]string, len(scoreBoard)), make(map[string]int)

	for i, item := range scoreBoard {
		scoreBoard[i].Progress = make([]responses.ScoreBoardProgress, contest.Count)
		uidList[i] = item.Uid
		mapUidToIndex[item.Uid] = i
	}

	var contestSubmissionList []models.ContestSubmission
	if err := db.DB.Table("contest_submission").
		Where("cid = ? AND uid IN (?)", cid, uidList).
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
			scoreBoard[mapUidToIndex[item.Uid]].Progress[questionId].Status = StatusAC

			totalTime := &scoreBoard[mapUidToIndex[item.Uid]].Progress[questionId].TotalTime
			if *totalTime == 0 || *totalTime > item.TotalTime {
				*totalTime = item.TotalTime
			}
		} else if item.Status == StatusERR {
			scoreBoard[mapUidToIndex[item.Uid]].Progress[questionId].Bug++
		}
	}

	// calc penalty
	for i := range scoreBoard {
		for j := range scoreBoard[i].Progress {
			if scoreBoard[i].Progress[j].Status == StatusAC {
				scoreBoard[i].Progress[j].TotalTime += scoreBoard[i].Progress[j].Bug * contest.Penalty
			}
		}
	}

	return scoreBoard, nil
}
