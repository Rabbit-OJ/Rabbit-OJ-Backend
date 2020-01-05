package models

import "time"

type ContestUser struct {
	Cuid      string    `gorm:"AUTO_INCREMENT" json:"cuid"`
	Cid       string    `json:"cid"`
	Uid       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	Score     uint32    `json:"score"`
	TotalTime uint32    `json:"total_time"`
}
