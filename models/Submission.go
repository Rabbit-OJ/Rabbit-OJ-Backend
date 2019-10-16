package models

import "time"

type Submission struct {
	Sid       string        `json:"sid"`
	Tid       string        `json:"tid"`
	Uid       string        `json:"uid"`
	Status    uint8         `json:"status"`
	Judge     []JudgeResult `json:"judge"` // TODO: CHECK JSON support
	Language  string        `json:"language"`
	FileName  string        `json:"-"`
	TimeUsed  uint32        `json:"time_used"`
	SpaceUsed uint32        `json:"space_used"`
	CreatedAt time.Time     `json:"created_at"`
}

type SubmissionLite struct {
	Sid       string    `json:"sid"`
	Uid       string    `json:"uid"`
	Status    uint8     `json:"status"`
	Language  string    `json:"language"`
	TimeUsed  uint32    `json:"time_used"`
	SpaceUsed uint32    `json:"space_used"`
	CreatedAt time.Time `json:"created_at"`
}

type JudgeResult struct {
	Status    uint8  `json:"status"`
	TimeUsed  uint32 `json:"time_used"`
	SpaceUsed uint32 `json:"space_used"`
}
