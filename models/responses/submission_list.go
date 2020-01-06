package responses

import "Rabbit-OJ-Backend/models"

type SubmissionListResponse struct {
	List  []models.SubmissionLite `json:"list"`
	Count int64                   `json:"count"`
}
