package question

import (
	"Rabbit-OJ-Backend/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAttemptCount(tid string) {
	if err := db.DB.Table("question").
		Update("attempt", gorm.Expr("attempt + 1")).
		Where("tid = ?", tid).Error;
		err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(tid string) {
	if err := db.DB.Table("question").
		Update("accept", gorm.Expr("accept + 1")).
		Where("tid = ?", tid).Error;
		err != nil {

		fmt.Println(err)
	}
}
