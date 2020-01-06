package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"time"
)

func Register(form *forms.RegisterForm) (uint32, error) {
	if UsernameExist(form.Username) {
		return InvalidUid, errors.New("username already exists")
	}

	if EmailExist(form.Email) {
		return InvalidUid, errors.New("email already exists")
	}

	user := models.User{
		Username:  form.Username,
		Password:  utils.SaltPasswordWithSecret(form.Password),
		Email:     form.Email,
		LoginAt:   time.Now(),
		CreatedAt: time.Now(),
		Attempt:   0,
		Accept:    0,
	}

	if _, err := db.DB.Table("user").Insert(&user); err != nil {
		return InvalidUid, err
	}

	return user.Uid, nil
}
