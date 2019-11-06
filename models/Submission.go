package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Submission struct {
	Sid           string        `json:"sid"`
	Tid           string        `json:"tid"`
	QuestionTitle string        `json:"question_title"`
	Uid           string        `json:"uid"`
	Status        string        `json:"status"`
	Judge         []byte        `json:"-"`
	JudgeObj      []JudgeResult `gorm:"-" json:"judge"`
	Language      string        `json:"language"`
	FileName      string        `json:"-"`
	TimeUsed      uint32        `json:"time_used"`
	SpaceUsed     uint32        `json:"space_used"`
	CreatedAt     time.Time     `json:"created_at"`
}

type SubmissionLite struct {
	Sid           string    `json:"sid"`
	Uid           string    `json:"uid"`
	Tid           string    `json:"tid"`
	QuestionTitle string    `json:"question_title"`
	Status        string    `json:"status"`
	Language      string    `json:"language"`
	TimeUsed      uint32    `json:"time_used"`
	SpaceUsed     uint32    `json:"space_used"`
	CreatedAt     time.Time `json:"created_at"`
}

type JudgeResult struct {
	Status    string `json:"status"`
	TimeUsed  uint32 `json:"time_used"`
	SpaceUsed uint32 `json:"space_used"`
}

func (s *Submission) AfterFind(scope *gorm.Scope) (err error) {
	s.JudgeObj = make([]JudgeResult, 0)
	if err := json.Unmarshal(s.Judge, &s.JudgeObj); err != nil {
		return err
	}
	return nil
}

func (s *Submission) BeforeSave(scope *gorm.Scope) (err error) {
	judgeJSON, err := json.Marshal(s.JudgeObj)
	if err != nil {
		return err
	}
	s.Judge = judgeJSON
	return nil
}
