package user

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAttemptCount(uid string) {
	if err := db.DB.Table("user").
		Where("uid = ?", uid).
		Update("attempt", gorm.Expr("attempt + 1")).
		Error;
		err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(uid string) {
	if err := db.DB.Table("user").
		Where("uid = ?", uid).
		Update("accept", gorm.Expr("accept + 1")).
		Error;
		err != nil {

		fmt.Println(err)
	}
}
