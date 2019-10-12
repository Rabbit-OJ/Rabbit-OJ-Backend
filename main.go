package main

import (
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	utils.InitConstant()
	server := gin.Default()

	server.Use(middlewares.Cors())
	server.GET("/user/login", user.Login)

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
