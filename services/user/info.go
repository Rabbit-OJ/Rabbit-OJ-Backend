package user

import (
	"Rabbit-OJ-Backend/db"
	model "Rabbit-OJ-Backend/models"
)

func InfoByUsername(username string) (*model.User, error) {
	user := model.User{}
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
