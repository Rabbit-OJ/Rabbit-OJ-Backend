package responses

type ScoreBoard struct {
	Uid       string               `json:"uid"`
	Username  string               `json:"username"`
	Score     uint32               `json:"score"`
	TotalTime uint32               `json:"total_time"`
	Progress  []ScoreBoardProgress `json:"progress" gorm:"-"`
}

type ScoreBoardProgress struct {
	Status    int    `json:"status"`
	Bug       uint32 `json:"bug"`
	TotalTime uint32 `json:"total_time"`
}
