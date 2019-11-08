package question

import (
	"Rabbit-OJ-Backend/db"
	"fmt"
)

func UpdateAttemptCount(tid string) {
	if err := db.DB.Table("question").Update("attempt = attempt + 1").Where("tid = ?", tid).Error;
		err != nil {
		fmt.Println(err)
	}
}

func UpdateAcceptedCount(tid string) {
	if err := db.DB.Table("question").Update("accepted = accepted + 1").Where("tid = ?", tid).Error;
		err != nil {
		fmt.Println(err)
	}
}