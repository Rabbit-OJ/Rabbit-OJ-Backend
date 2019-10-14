package models

import (
	"time"
)

type User struct {
	Uid       uint32 `gorm:"auto_increment"`
	Username  string
	IsAdmin   bool
	Password  string
	Email     string
	Attempt   uint32 `gorm:"default:0'"`
	Accept    uint32 `gorm:"default:0'"`
	LoginAt   time.Time
	CreatedAt time.Time
}

type OtherUser struct {
	Uid       uint32
	Username  string
	IsAdmin   bool
	Attempt   uint32
	Accept    uint32
	LoginAt   time.Time
	CreatedAt time.Time
}

type MyUser struct {
	Uid       uint32
	Username  string
	IsAdmin   bool
	Email     string
	Attempt   uint32
	Accept    uint32
	LoginAt   time.Time
	CreatedAt time.Time
}