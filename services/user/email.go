package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
)

const InvalidUid = ""

func EmailToUid(email string) (string, error) {
	user := models.User{}

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return InvalidUid, err
	}

	return user.Uid, nil
}

func EmailExist(email string) bool {
	uid, _ := EmailToUid(email)

	return uid != InvalidUid
}