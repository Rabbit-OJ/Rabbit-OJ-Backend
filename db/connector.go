package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	DB *gorm.DB
)

func Init() {
	username := "root"
	password := "P@ssw0rd"
	database := "rabbit"

	server := "tcp(localhost:3306)"
	if os.Getenv("ENV") == "production" {
		server = "tcp(mysql:3306)"
	}

	connStr := fmt.Sprintf("%s:%s@%s/%s?&parseTime=True&loc=Local", username, password, server, database)
	db, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	//db.LogMode(true)
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	DB = db
}
