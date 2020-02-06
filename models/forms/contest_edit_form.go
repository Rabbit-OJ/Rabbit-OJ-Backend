package forms

type ContestEditForm struct {
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	BlockTime int64  `json:"block_time"`
	EndTime   int64  `json:"end_time"`
	Penalty   uint32 `json:"penalty"`
	Status    int32  `json:"status"`
}
