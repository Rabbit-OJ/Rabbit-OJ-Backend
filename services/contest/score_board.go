package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
	"strconv"
	"time"
)

func InBlockTime(cid uint32) (bool, error) {
	contest, err := Info(cid)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()
	return time.Time(contest.BlockTime).Unix() <= now && now <= time.Time(contest.EndTime).Unix(), nil
}

func ScoreBoard(contest *models.Contest, page uint32) ([]*responses.ScoreBoard, bool, error) {
	scoreBoard := make([]*responses.ScoreBoard, 0)
	cid := contest.Cid

	inBlockTime, err := InBlockTime(cid)
	if err != nil {
		return nil, false, err
	}

	if inBlockTime {
		return nil, true, nil
	}

	questionMapTidToId, _, err := QuestionMapTidToId(cid)
	if err != nil {
		return nil, false, err
	}

	results, err := db.DB.QueryString("SELECT `user`.username, `user`.uid, `rank`, score, total_time FROM ( "+
		"SELECT cid, uid, score, total_time, RANK() OVER (ORDER BY score DESC, total_time ASC) `rank` FROM ( "+
		"SELECT cid, uid, score, total_time FROM contest_user WHERE cid = ? "+
		") AS temp1 "+
		") AS temp2 "+
		"INNER JOIN `user` ON `temp2`.`uid`=`user`.`uid` "+
		"LIMIT ?, ? ",
		cid, (page-1)*config.PageSize, page*config.PageSize)

	for _, item := range results {
		_score, _ := strconv.ParseUint(item["score"], 10, 32)
		_totalTime, _ := strconv.ParseUint(item["total_time"], 10, 32)
		_rank, _ := strconv.ParseUint(item["rank"], 10, 32)
		_uid, _ := strconv.ParseUint(item["uid"], 10, 32)

		scoreBoard = append(scoreBoard, &responses.ScoreBoard{
			Uid:       uint32(_uid),
			Username:  item["username"],
			Score:     uint32(_score),
			TotalTime: uint32(_totalTime),
			Rank:      uint32(_rank),
			Progress:  nil,
		})
	}

	if err != nil {
		return nil, false, err
	}

	uidList, mapUidToIndex := make([]uint32, len(scoreBoard)), make(map[uint32]int)
	for i, item := range scoreBoard {
		scoreBoard[i].Progress = make([]responses.ScoreBoardProgress, contest.Count)
		uidList[i] = item.Uid
		mapUidToIndex[item.Uid] = i
	}

	var contestSubmissionList []models.ContestSubmission
	if err := db.DB.Table("contest_submission").
		Where("cid = ?", cid).
		In("uid", uidList).
		Find(&contestSubmissionList); err != nil {
		return nil,false,  err
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

	return scoreBoard,false,  nil
}
