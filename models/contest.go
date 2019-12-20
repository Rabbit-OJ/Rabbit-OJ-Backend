package models

import "time"

type Contest struct {
	Cid          string    `gorm:"auto_increment" json:"cid"`
	Name         string    `json:"name"`
	Uid          string    `json:"uid"`
	Count        uint32    `json:"count"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	BlockTime    time.Time `json:"block_time"`
	Status       int32     `json:"status"`
	Participants uint32    `json:"participants"`
	Penalty      uint32    `json:"penalty"`
}
