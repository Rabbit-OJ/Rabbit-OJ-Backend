package models

import "time"

type ContestUser struct {
	Cuid      uint32    `xorm:"autoincr" json:"cuid"`
	Cid       uint32    `json:"cid"`
	Uid       uint32    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	Score     uint32    `json:"score"`
	TotalTime uint32    `json:"total_time"`
}
