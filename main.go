package main

import (
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/services/initialize"
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/services/routine"
	"Rabbit-OJ-Backend/services/rpc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	Role = ""
)

func main() {
	globalContext, cancelGlobalContext := context.WithCancel(context.Background())
	defer func() {
		cancelGlobalContext()
	}()

	Role = os.Getenv("Role")
	fmt.Printf("[Role] %s \n", Role)

	if Role == "Server" {
		initialize.Config()

		routine.StartCheck()
		initialize.Cert("server")
		initialize.DB(globalContext)
		initialize.MQ(globalContext)

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

		initialize.Cert("client")
		initialize.DindScript()
		initialize.Docker()
		initialize.CheckTestCase()

		initialize.MQ(globalContext)
		routine.RegisterSignal()

		exitChan := make(chan bool)

		select {
		case <-exitChan:
			os.Exit(0)
		}
	} else if Role == "Tester" {
		judger.Tester()
	}
}
