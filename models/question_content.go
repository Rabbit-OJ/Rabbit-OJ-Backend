package models

type QuestionContent struct {
	Tid     uint32   `json:"tid"`
	Content string   `json:"content"`
	Sample  []Sample `json:"sample"`
}

type Sample struct {
	In  string `json:"in"`
	Out string `json:"out"`
}
