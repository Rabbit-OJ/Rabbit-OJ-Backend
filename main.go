package main

import (
	"Rabbit-OJ-Backend/controllers/question"
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

	server.GET("/user/info/:username", user.Info)
	server.GET("/user/my", middlewares.AuthJWT(), user.My)
	server.POST("/user/login", user.Login)
	server.POST("/user/register", user.Register)

	server.GET("/submission/list/:uid/:page", submission.List)
	server.GET("/submission/detail/:sid", submission.Detail)

	server.GET("/question/list/:page", question.List)
	server.GET("/question/content/:tid", question.Detail)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
