package forms

type ContestQuestionEditForm struct {
	Data []ContestQuestionEditItem `json:"data" binding:"required"`
}

type ContestQuestionEditItem struct {
	Tid   uint32 `json:"tid"`
	Id    int    `json:"id"`
	Score uint32 `json:"score"`
}
