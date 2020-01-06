package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

const InvalidUid = 99

func EmailToUid(email string) (uint32, error) {
	user := models.User{}

	found, err := db.DB.Table("user").Where("email = ?", email).Get(&user)
	if err != nil {
		return InvalidUid, err
	}
	if !found {
		return InvalidUid, errors.New("user doesn't exist")
	}

	return user.Uid, nil
}

func EmailExist(email string) bool {
	uid, _ := EmailToUid(email)

	return uid != InvalidUid
}
