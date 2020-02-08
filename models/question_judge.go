package models

type QuestionJudge struct {
	Tid          uint32 `json:"tid"`
	Mode         string `json:"mode"`
	DatasetCount uint32 `json:"dataset_count"`
	Version      uint32 `json:"version"`
}
