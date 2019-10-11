package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uid string `gorm:"AUTO_INCREMENT"`
	Password string
	Email string
	Attempt int32
	Accepted int32
	RegisterDate time.Time
	LastLoginDate time.Time
}