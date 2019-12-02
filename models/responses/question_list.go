package responses

import "Rabbit-OJ-Backend/models"

type QuestionList struct {
	List  []models.Question `json:"list"`
	Count uint32            `json:"count"`
}
