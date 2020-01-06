package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"errors"
)

func MyInfoByUid(uid uint32) (*models.MyUser, error) {
	user := models.MyUser{}

	found, err := db.DB.Table("user").Where("uid = ?", uid).Get(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user doesn't exist")
	}

	return &user, nil
}

func InfoByUsername(username string) (*models.User, error) {
	user := models.User{}

	found, err := db.DB.Table("user").Where("username = ?", username).Get(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user doesn't exist")
	}

	return &user, nil
}

func InfoByUid(uid uint32) (*models.User, error) {
	user := models.User{}

	found, err := db.DB.Table("user").Where("uid = ?", uid).Get(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user doesn't exist")
	}

	return &user, nil
}

func OtherInfoByUsername(username string) (*models.OtherUser, error) {
	user := models.OtherUser{}

	found, err := db.DB.Table("user").Where("username = ?", username).Get(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("user doesn't exist")
	}

	return &user, nil
}
