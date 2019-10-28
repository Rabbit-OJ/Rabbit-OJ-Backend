package user

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func MyInfoByUid(uid string) (*models.MyUser, error) {
	user := models.MyUser{}
	if err := db.DB.Table("user").Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func InfoByUsername(username string) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func InfoByUid(uid string) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func OtherInfoByUsername(username string) (*models.OtherUser, error) {
	user := models.OtherUser{}
	if err := db.DB.Table("user").Where("username = ?", username).Scan(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}