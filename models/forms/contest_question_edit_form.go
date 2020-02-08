package forms

type ContestQuestionEditForm struct {
	Data []ContestQuestionEditItem `json:"data" binding:"required"`
}

type ContestQuestionEditItem struct {
	Tid   uint32 `json:"tid" binding:"required"`
	Id    int    `json:"id" binding:"required"`
	Score uint32 `json:"score" binding:"required"`
}
