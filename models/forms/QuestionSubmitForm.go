package forms

type QuestionSubmitForm struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code" binding:"code"`
}
