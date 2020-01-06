package responses

import "Rabbit-OJ-Backend/models"

type QuestionListResponse struct {
	List  []models.Question `json:"list"`
	Count int64             `json:"count"`
}
