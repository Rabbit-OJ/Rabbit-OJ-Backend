package question

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAttemptCount(tid string) {
	if err := db.DB.Table("question").
		Where("tid = ?", tid).
		Update("attempt", gorm.Expr("attempt + 1")).
		Error;
		err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(tid string) {
	if err := db.DB.Table("question").
		Where("tid = ?", tid).
		Update("accept", gorm.Expr("accept + 1")).
		Error;
		err != nil {

		fmt.Println(err)
	}
}
