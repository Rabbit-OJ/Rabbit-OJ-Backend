package user

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func EmailToUid(email string) (uint32, error) {
	user := models.User{}

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, err
	}

	return user.Uid, nil
}

func EmailExist(email string) bool {
	uid, _ := EmailToUid(email)

	return uid != 0
}