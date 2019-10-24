package main

import (
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/mq"
	"Rabbit-OJ-Backend/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	mq.Init()
	utils.InitConstant()

	defer func() {
		if err := db.DB.Close(); err != nil {
			fmt.Println(err)
		}

		if err := mq.Connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	server := gin.Default()

	server.Use(middlewares.Cors())
	user.Router(server)
	submission.Router(server)
	question.Router(server)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
