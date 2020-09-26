package main

import (
	"Rabbit-OJ-Backend/controllers/contest"
	"Rabbit-OJ-Backend/controllers/question"
	"Rabbit-OJ-Backend/controllers/submission"
	"Rabbit-OJ-Backend/controllers/user"
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/services/initialize"
	"Rabbit-OJ-Backend/services/routine"
	"Rabbit-OJ-Backend/services/storage/rpc"
	"context"
	"fmt"
	"github.com/Rabbit-OJ/Rabbit-OJ-Tester"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Role = ""
)

func main() {
	config.GlobalContext, config.CancelGlobalContext = context.WithCancel(context.Background())
	defer func() {
		config.CancelGlobalContext()
	}()

	Role = os.Getenv("Role")
	fmt.Printf("[Role] %s \n", Role)

	if Role == "Server" {
		initialize.Config()

		routine.StartCheck()
		initialize.Cert("server")
		initialize.Judger(config.GlobalContext)
		initialize.DB(config.GlobalContext)

		go rpc.Register()
		server := gin.Default()

		server.Use(middlewares.Cors())
		user.Router(server)
		submission.Router(server)
		question.Router(server)
		contest.Router(server)
		websocket.WebSocket(server)

		err := server.Run(":8888")
		if err != nil {
			fmt.Println(err)
		}
	} else if Role == "Judge" {
		initialize.Config()

		initialize.Cert("client")
		initialize.Judger(config.GlobalContext)
		initialize.CheckTestCase()

		routine.RegisterSignal()

		exitChan := make(chan bool)

		select {
		case <-exitChan:
			os.Exit(0)
		}
	} else if Role == "Tester" {
		fmt.Println("[WARN] use v1.0 Tester is deprecated, and will be removed soon, please use v2 Tester instead.")
		tester.Tester()
	}
}
