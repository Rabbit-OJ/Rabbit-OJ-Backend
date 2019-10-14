package user

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"time"
)

func Register(form *forms.RegisterForm) (uint32, error) {
	if UsernameExist(form.Username) {
		return 0, errors.New("username already exists")
	}

	if EmailExist(form.Email) {
		return 0, errors.New("email already exists")
	}

	user := models.User{
		Username: form.Username,
		Password: utils.SaltPasswordWithSecret(form.Password),
		Email:    form.Email,
		LoginAt:  time.Now(),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.Uid, nil
}
