package main

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/initialize"
	"context"
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	initialize.Config()
	initialize.DB(context.Background())

	if _, err := db.DB.Table("contest_submission").
		In("sid", []uint32{18, 19}).
		Update(
			&models.ContestSubmission{
				Status: -1,
			}); err != nil {

		fmt.Println(err)
	}
}
