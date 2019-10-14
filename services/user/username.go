package user

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func UsernameToUid(username string) (uint32, error) {
	user := models.User{}

	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}

	return user.Uid, nil
}

func UsernameExist(username string) bool {
	uid, _ := UsernameToUid(username)

	return uid != 0
}