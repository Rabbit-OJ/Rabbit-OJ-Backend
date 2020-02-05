package forms

import "Rabbit-OJ-Backend/models"

type QuestionForm struct {
	Subject    string          `json:"subject" binding:"required"`
	Difficulty uint8           `json:"difficulty" binding:"required"`
	TimeLimit  uint32          `json:"time_limit" binding:"required"`
	SpaceLimit uint32          `json:"space_limit" binding:"required"`
	Content    string          `json:"content" binding:"required"`
	Hide       bool            `json:"hide" binding:"required"`
	Sample     []models.Sample `json:"sample" binding:"required"`
}
