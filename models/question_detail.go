package models

import (
	"time"
)

type QuestionDetail struct {
	Tid        string    `json:"tid"`
	Content    string    `json:"content"`
	Sample     []Sample  `json:"sample"`
	Subject    string    `json:"subject"`
	Hide       bool      `json:"hide"`
	Attempt    uint32    `json:"attempt"`
	Accept     uint32    `json:"accept"`
	Difficulty uint8     `json:"difficulty"`
	TimeLimit  uint32    `json:"time_limit"`
	SpaceLimit uint32    `json:"space_limit"`
	CreatedAt  time.Time `json:"created_at"`
}
