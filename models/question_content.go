package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type QuestionContent struct {
	Tid        string   `json:"tid"`
	Content    string   `json:"content"`
	Sample     []byte   `json:"-"`
	SampleJSON []Sample `json:"Sample",gorm:"-"`
}

type Sample struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

func (s *QuestionContent) AfterFind(_ *gorm.Scope) (err error) {
	s.SampleJSON = make([]Sample, 0)
	if err := json.Unmarshal(s.Sample, &s.SampleJSON); err != nil {
		return err
	}
	return nil
}

func (s *QuestionContent) BeforeSave(_ *gorm.Scope) (err error) {
	sampleJSON, err := json.Marshal(s.SampleJSON)
	if err != nil {
		return err
	}
	s.Sample = sampleJSON
	return nil
}
