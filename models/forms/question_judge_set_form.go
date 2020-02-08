package forms

type QuestionJudgeSetForm struct {
	Mode         string `json:"mode" binding:"required"`
	DatasetCount uint32 `json:"dataset_count" binding:"required"`
	Version      uint32 `json:"version"`
}
