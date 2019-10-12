package main

import (
	"Rabbit-OJ-Backend/controllers/User"
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	server := gin.Default()

	server.Use(middlewares.Cors())
	server.GET("/login", User.Login)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		err := db.DB.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
