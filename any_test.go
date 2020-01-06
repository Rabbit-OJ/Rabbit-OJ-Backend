package main

import (
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/services/initialize"
	"context"
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	initialize.Config()
	initialize.DB(context.Background())

	_, err := db.DB.Table("user").
		Where("uid = ?", "1").
		Update("attempt = attempt + 1")

	if err != nil {
		fmt.Println(err)
	}
}
