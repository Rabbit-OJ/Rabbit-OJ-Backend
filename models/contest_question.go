package models

import (
	"time"
)

type ContestQuestion struct {
	Cid   string `json:"cid"`
	Tid   string `json:"tid"`
	Id    int    `json:"id"`
	Score uint32 `json:"score"`
}

type ContestQuestionExtended struct {
	Cid        string    `json:"cid"`
	Tid        string    `json:"tid"`
	Id         int       `json:"id"`
	Score      uint32    `json:"score"`
	Uid        string    `json:"uid"`
	Subject    string    `json:"subject"`
	Difficulty uint8     `json:"difficulty"`
	TimeLimit  uint32    `json:"time_limit"`
	SpaceLimit uint32    `json:"space_limit"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	Sample     []byte    `json:"-"`
	SampleJSON []Sample  `json:"sample",gorm:"-"`
}