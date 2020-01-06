package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/db"
	"context"
	"fmt"
	"xorm.io/xorm"
)

func DB(globalContext context.Context) {
	connStr := config.Global.MySQL
	conn, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic(err)
	}

	db.DB = conn

	go func() {
		<-globalContext.Done()
		if err := db.DB.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}
