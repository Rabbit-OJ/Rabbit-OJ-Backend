package responses

type ContestMyInfo struct {
	Rank      uint32               `json:"rank"`
	TotalTime uint32               `json:"total_time"`
	Score     uint32               `json:"score"`
	Progress  []ScoreBoardProgress `json:"progress"`
}
