package question

import (
	"Rabbit-OJ-Backend/services/db"
	"fmt"
)

func UpdateAttemptCount(tid uint32) {
	if _, err := db.DB.Exec("UPDATE question SET attempt = attempt + 1 WHERE tid = ?", tid); err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(tid uint32) {
	if _, err := db.DB.Exec("UPDATE question SET accept = accept + 1 WHERE tid = ?", tid); err != nil {

		fmt.Println(err)
	}
}
