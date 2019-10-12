package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Uid string `gorm:"AUTO_INCREMENT"`
	Username string
	Password string
	Email string
	Attempt int32
	Accepted int32
	loginAt time.Time
	CreatedAt time.Time
}