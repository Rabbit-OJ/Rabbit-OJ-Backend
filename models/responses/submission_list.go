package responses

import "Rabbit-OJ-Backend/models"

type SubmissionList struct {
	List  []models.SubmissionLite `json:"list"`
	Count uint32                  `json:"count"`
}
