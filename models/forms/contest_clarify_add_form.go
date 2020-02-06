package forms

type ContestClarifyAddForm struct {
	Cid     uint32 `json:"cid" binding:"required"`
	Message string `json:"message" binding:"required"`
}
