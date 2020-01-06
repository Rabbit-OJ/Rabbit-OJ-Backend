package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func UsernameToUid(username string) (uint32, error) {
	user := models.User{}

	found, err := db.DB.Where("username = ?", username).Get(&user)
	if err != nil {
		return InvalidUid, err
	}
	if !found {
		return InvalidUid, errors.New("user doesn't exist")
	}

	return user.Uid, nil
}

func UsernameExist(username string) bool {
	uid, _ := UsernameToUid(username)

	return uid != InvalidUid
}
