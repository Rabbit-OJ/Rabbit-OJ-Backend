package main

import (
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer func() {
		err := db.DB.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	utils.InitConstant()
	server := gin.Default()

	server.Use(middlewares.Cors())

	server.GET("/user/login", user.Login)
	server.GET("/user/info/:username", user.Info)
	server.GET("/user/my", user.My)
	server.POST("/user/register", user.Register)
	server.GET("/submission/:uid/:page", submission.List)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
}
