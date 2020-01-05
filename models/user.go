package models

import (
	"time"
)

type User struct {
	Uid       string    `gorm:"AUTO_INCREMENT" json:"uid"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"isAdmin"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Attempt   uint32    `gorm:"default:0'" json:"attempt"`
	Accept    uint32    `gorm:"default:0'" json:"accept"`
	LoginAt   time.Time `json:"loginAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type OtherUser struct {
	Uid       string    `json:"uid"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"isAdmin"`
	Attempt   uint32    `json:"attempt"`
	Accept    uint32    `json:"accept"`
	LoginAt   time.Time `json:"loginAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type MyUser struct {
	Uid       string    `json:"uid"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"isAdmin"`
	Email     string    `json:"email"`
	Attempt   uint32    `json:"attempt"`
	Accept    uint32    `json:"accept"`
	LoginAt   time.Time `json:"loginAt"`
	CreatedAt time.Time `json:"createdAt"`
}
