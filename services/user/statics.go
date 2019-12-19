package user

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAttemptCount(uid string) {
	if err := db.DB.Table("user").
		Update("attempt", gorm.Expr("attempt + 1")).
		Where("uid = ?", uid).Error;
		err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(uid string) {
	if err := db.DB.Table("user").
		Update("accept", gorm.Expr("accept + 1")).
		Where("uid = ?", uid).Error;
		err != nil {

		fmt.Println(err)
	}
}
