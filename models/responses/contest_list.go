package responses

import "Rabbit-OJ-Backend/models"

type ContestList struct {
	List  []models.Contest `json:"list"`
	Count int64            `json:"count"`
}
