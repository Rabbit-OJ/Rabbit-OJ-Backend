package main

import (
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/controllers/websocket"
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

var (
	Role = ""
)

func main() {
	Role = os.Getenv("Role")
	utils.InitConstant()

	if Role == "Server" {
		db.Init()
		mq.Init()
		judger.InitMQ()

		defer func() {
			if err := mq.Channel.Close(); err != nil {
				fmt.Println(err)
			}

			if err := mq.Connection.Close(); err != nil {
				fmt.Println(err)
			}

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
		websocket.WebSocket(server)

		err := server.Run(":8888")
		if err != nil {
			fmt.Println(err)
		}
	} else if Role == "Judge" {
		mq.Init()
		judger.InitMQ()

		defer func() {
			if err := mq.Channel.Close(); err != nil {
				fmt.Println(err)
			}

			if err := mq.Connection.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		judger.InitDocker()
		judger.CheckTestCase()

		exitChan := make(chan bool)

		select {
		case <-exitChan:
			os.Exit(0)
		}
	} else if Role == "Tester" {
		judger.Tester()
	}
}
