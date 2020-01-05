package models

import "time"

type ContestClarify struct {
	Ccid      string    `gorm:"AUTO_INCREMENT" json:"ccid"`
	Cid       string    `json:"cid"`
	CreatedAt time.Time `json:"created_at"`
	Message string `json:"message"`
}
