package user

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	userRouter := baseRouter.Group("/user")
	userRouter.GET("/info/:username", Info)
	userRouter.GET("/my", middlewares.AuthJWT(), My)
	userRouter.POST("/login", Login)
	userRouter.POST("/register", Register)
}
