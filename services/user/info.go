package user

import (
	"Rabbit-OJ-Backend/db"
	model "Rabbit-OJ-Backend/models"
)

func InfoByUsername(username string) (user *model.User) {
	db.DB.Where("username = ?", username).First(&user)
	return
}
