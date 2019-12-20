package models

import "time"

type ContestUser struct {
	Cuid      string    `gorm:"auto_increment" json:"cuid"`
	Cid       string    `json:"cid"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	Rank      uint32    `json:"rank"`
	Score     uint32    `json:"score"`
	TotalTime uint32    `json:"total_time"`
}
