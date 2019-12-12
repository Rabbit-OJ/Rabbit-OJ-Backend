package main

import (
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/services/initialize"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	Role = ""
)

func main() {
	exitChan := make(chan bool)
	defer func() {
		exitChan <- true
		close(exitChan)
	}()

	Role = os.Getenv("Role")

	if Role == "Server" {
		initialize.Config()
		initialize.DB(exitChan)
		initialize.MQ(exitChan)

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
		initialize.Config()

		initialize.DindScript()
		initialize.Docker()
		initialize.CheckTestCase()

		initialize.MQ(exitChan)

		exitChan := make(chan bool)

		select {
		case <-exitChan:
			os.Exit(0)
		}
	} else if Role == "Tester" {
		judger.Tester()
	}
}
