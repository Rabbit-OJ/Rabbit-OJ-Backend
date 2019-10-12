package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init() {
	username := "root"
	password := "P@ssw0rd"
	database := "rabbit"

	connStr := fmt.Sprintf("%s:%s@/%s?&parseTime=True&loc=Local", username, password, database)
	db, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	DB = db
}
