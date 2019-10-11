package main

import (
	"Rabbit-OJ-Backend/controller/User"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// server map
	server.GET("/login", User.Login)

	// err handler
	err := server.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

}
