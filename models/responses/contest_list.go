package responses

import "Rabbit-OJ-Backend/models"

type ContestList struct {
	List  []models.Contest `json:"list"`
	Count uint32           `json:"count"`
}
