package models

import "time"

type ContestSubmission struct {
	Sid       string    `json:"csid"`
	Cid       string    `json:"cid"`
	Uid       string    `json:"uid"`
	Tid       string    `json:"tid"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	TotalTime uint32    `json:"total_time"`
}
