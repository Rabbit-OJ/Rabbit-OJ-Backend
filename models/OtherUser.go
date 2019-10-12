package models

import "time"

type OtherUser struct {
	Uid       string `gorm:"AUTO_INCREMENT"`
	Username  string
	IsAdmin   bool
	Attempt   uint32
	Accept    uint32
	LoginAt   time.Time
	CreatedAt time.Time
}