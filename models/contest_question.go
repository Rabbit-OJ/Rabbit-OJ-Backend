package models

type ContestQuestion struct {
	Cid    string `json:"cid"`
	Tid    string `json:"tid"`
	Id     int    `json:"id"`
	Score  uint32 `json:"score"`
}
