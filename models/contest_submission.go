package models

import "time"

type ContestSubmission struct {
	Sid       string    `json:"sid"`
	Cid       string    `json:"cid"`
	Uid       string    `json:"uid"`
	Tid       string    `json:"tid"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	TotalTime int64     `json:"total_time"`
}
