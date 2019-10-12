package user

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/utils"
	"time"
)

func Register(form *models.RegisterForm) (string, error) {
	user := models.User{
		Username: form.Username,
		Password: utils.SaltPasswordWithSecret(form.Password),
		Email:    form.Email,
		Accept:   0,
		Attempt:  0,
		LoginAt:  time.Now(),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return "", err
	}

	return user.Uid, nil
}
