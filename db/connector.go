package db

import (
	"Rabbit-OJ-Backend/services/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init() {
	connStr := config.Global.MySQL
	db, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	DB = db
}
