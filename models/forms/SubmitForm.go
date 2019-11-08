package forms

type SubmitForm struct {
	Language string `json:"language" binding:"required,language"`
	Code     string `json:"code" binding:"required"`
}
