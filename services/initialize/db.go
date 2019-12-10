package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func DB(exitChan chan bool) {
	connStr := config.Global.MySQL
	conn, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	conn.SingularTable(true)
	db.DB = conn

	go func() {
		<-exitChan
		if err := db.DB.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}
