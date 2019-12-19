package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type QuestionDetail struct {
	Tid        string    `json:"tid"`
	Content    string    `json:"content"`
	Sample     []byte    `json:"-"`
	SampleJSON []sample  `json:"sample",gorm:"-"`
	Subject    string    `json:"subject"`
	Hide       bool      `json:"hide"`
	Attempt    uint32    `json:"attempt"`
	Accept     uint32    `json:"accept"`
	Difficulty uint8     `json:"difficulty"`
	TimeLimit  uint32    `json:"time_limit"`
	SpaceLimit uint32    `json:"space_limit"`
	CreatedAt  time.Time `json:"created_at"`
}

func (s *QuestionDetail) AfterFind(_ *gorm.Scope) (err error) {
	s.SampleJSON = make([]sample, 0)
	if err := json.Unmarshal(s.Sample, &s.SampleJSON); err != nil {
		return err
	}
	return nil
}

func (s *QuestionDetail) BeforeSave(_ *gorm.Scope) (err error) {
	sampleJSON, err := json.Marshal(s.SampleJSON)
	if err != nil {
		return err
	}
	s.Sample = sampleJSON
	return nil
}
