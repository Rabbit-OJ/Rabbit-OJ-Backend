package main

import (
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/mq"
	"Rabbit-OJ-Backend/services/rpc"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	mq.Init()
	utils.InitConstant()

	defer func() {
		if err := mq.Channel.Close(); err != nil {
			fmt.Println(err)
		}

		if err := mq.Connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if os.Getenv("Role") == "Server" {
		db.Init()

		defer func() {
			if err := db.DB.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		go rpc.Register()
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

	if os.Getenv("Role") == "Judge" {
		judger.InitDocker()
		judger.CheckTestCase()
	}

	if os.Getenv("Role") == "Tester" {
		judger.Tester()
	}
}
