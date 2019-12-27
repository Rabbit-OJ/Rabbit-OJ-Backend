package responses

import "Rabbit-OJ-Backend/models"

type QuestionListResponse struct {
	List  []models.Question `json:"list"`
	Count uint32            `json:"count"`
}
