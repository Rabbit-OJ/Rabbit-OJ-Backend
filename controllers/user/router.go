package user

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	userRouter := baseRouter.Group("/user")

	userRouter.GET("/info/:username", Info)
	userRouter.GET("/avatar/:uid", Avatar)
	userRouter.GET("/my", middlewares.AuthJWT(), My)
	userRouter.POST("/my/avatar", middlewares.AuthJWT(), UploadAvatar)
	userRouter.POST("/login", Login)
	userRouter.POST("/register", Register)
}
