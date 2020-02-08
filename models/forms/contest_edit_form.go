package forms

type ContestEditForm struct {
	Name      string `json:"name" binding:"required"`
	StartTime int64  `json:"start_time" binding:"required"`
	BlockTime int64  `json:"block_time" binding:"required"`
	EndTime   int64  `json:"end_time" binding:"required"`
	Penalty   uint32 `json:"penalty" binding:"required"`
	Status    int32  `json:"status" binding:"required"`
}
