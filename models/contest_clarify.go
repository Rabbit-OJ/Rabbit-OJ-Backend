package models

import "time"

type ContestClarify struct {
	Ccid      uint32    `xorm:"autoincr" json:"ccid"`
	Cid       uint32    `json:"cid"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}
