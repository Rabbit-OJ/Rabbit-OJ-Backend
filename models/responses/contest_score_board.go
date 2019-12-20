package responses

type ScoreBoard struct {
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Score     uint32 `json:"score"`
	TotalTime uint32 `json:"total_time"`
}
