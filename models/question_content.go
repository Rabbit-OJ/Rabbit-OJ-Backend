package models

type QuestionContent struct {
	Tid        uint32   `json:"tid"`
	Content    string   `json:"content"`
	Sample     []byte   `json:"-"`
	SampleJSON []Sample `json:"sample",xorm:"-"`
}

type Sample struct {
	In  string `json:"in"`
	Out string `json:"out"`
}
