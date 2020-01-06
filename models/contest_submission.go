package models

import "time"

type ContestSubmission struct {
	Sid       uint32    `json:"sid"`
	Cid       uint32    `json:"cid"`
	Uid       uint32    `json:"uid"`
	Tid       uint32    `json:"tid"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	TotalTime uint32    `json:"total_time"`
}
