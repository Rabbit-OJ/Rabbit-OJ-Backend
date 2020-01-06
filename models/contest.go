package models

import (
	"fmt"
	"time"
)

type JSONTime time.Time

type Contest struct {
	Cid          uint32   `xorm:"autoincr" json:"cid"`
	Name         string   `json:"name"`
	Uid          string   `json:"uid"`
	Count        uint32   `json:"count"`
	StartTime    JSONTime `json:"start_time"`
	EndTime      JSONTime `json:"end_time"`
	BlockTime    JSONTime `json:"block_time"`
	Status       int32    `json:"status"`
	Participants uint32   `json:"participants"`
	Penalty      uint32   `json:"penalty"`
}

func (j *JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*j).Format("2006/01/02 15:04:05 -0700"))
	return []byte(stamp), nil
}
