package responses

type ScoreBoard struct {
	Uid       string               `json:"uid"`
	Username  string               `json:"username"`
	Score     uint32               `json:"score"`
	TotalTime uint32               `json:"total_time"`
	Rank      uint32               `json:"rank"`
	Progress  []ScoreBoardProgress `json:"progress" gorm:"-"`
}

type ScoreBoardProgress struct {
	Status    int    `json:"status"`
	Bug       uint32 `json:"bug"`
	TotalTime int64  `json:"total_time"`
}

type ScoreBoardResponse struct {
	List  []ScoreBoard `json:"list"`
	Count uint32       `json:"count"`
}
