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
	user.Router(server)
	submission.Router(server)
	question.Router(server)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
