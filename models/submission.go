package models

import (
	"time"
)

type Submission struct {
	Sid       uint32        `xorm:"autoincr" json:"sid"`
	Tid       uint32        `json:"tid"`
	Uid       uint32        `json:"uid"`
	Status    string        `json:"status"`
	Judge     []JudgeResult `json:"judge"`
	Language  string        `json:"language"`
	FileName  string        `json:"-"`
	TimeUsed  uint32        `json:"time_used"`
	SpaceUsed uint32        `json:"space_used"`
	CreatedAt time.Time     `json:"created_at"`
}

type SubmissionExtended struct {
	Sid           uint32        `json:"sid"`
	Tid           uint32        `json:"tid"`
	QuestionTitle string        `json:"question_title"`
	Uid           uint32        `json:"uid"`
	Status        string        `json:"status"`
	Judge         []JudgeResult `json:"judge"`
	Language      string        `json:"language"`
	FileName      string        `json:"-"`
	TimeUsed      uint32        `json:"time_used"`
	SpaceUsed     uint32        `json:"space_used"`
	CreatedAt     time.Time     `json:"created_at"`
}

type SubmissionLite struct {
	Sid           uint32    `json:"sid"`
	Uid           uint32    `json:"uid"`
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