package forms

type ContestClarifyAdd struct {
	Cid     uint32 `json:"cid" binding:"required"`
	Message string `json:"message" binding:"required"`
}
