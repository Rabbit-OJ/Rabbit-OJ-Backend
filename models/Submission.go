package models

import "time"

type Submission struct {
	Sid       uint32
	Uid       uint32
	Status    uint8
	Judge     []JudgeResult // TODO: CHECK JSON support
	Language  string
	FileName  string
	TimeUsed  uint32
	SpaceUsed uint32
	CreatedAt time.Time
}

type SubmissionLite struct {
	Sid       uint32
	Uid       uint32
	Status    uint8
	Language  string
	TimeUsed  uint32
	SpaceUsed uint32
	CreatedAt time.Time
}

type JudgeResult struct {
	Status    uint8  `json:"status"`
	TimeUsed  uint32 `json:"time_used"`
	SpaceUsed uint32 `json:"space_used"`
}
