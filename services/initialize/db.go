package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

func DB(globalContext context.Context) {
	connStr := config.Global.MySQL
	conn, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	conn.SingularTable(true)
	db.DB = conn

	go func() {
		<-globalContext.Done()
		if err := db.DB.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}
