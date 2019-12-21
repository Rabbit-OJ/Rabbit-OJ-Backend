package models

type ContestQuestion struct {
	Cid    string `json:"cid"`
	Tid    string `json:"tid"`
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Score  uint32 `json:"score"`
}
