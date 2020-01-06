package user

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
)

func UpdateAttemptCount(uid uint32) {
	if _, err := db.DB.Exec("UPDATE user SET attempt = attempt + 1 WHERE uid = ?", uid); err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(uid uint32) {
	if _, err := db.DB.Exec("UPDATE user SET accept = accept + 1 WHERE uid = ?", uid); err != nil {

		fmt.Println(err)
	}
}
