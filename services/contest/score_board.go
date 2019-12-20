package contest

import "time"

func InBlockTime(cid string) (bool, error) {
	contest, err := Info(cid)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()
	return contest.BlockTime.Unix() <= now && now <= contest.EndTime.Unix(), nil
}