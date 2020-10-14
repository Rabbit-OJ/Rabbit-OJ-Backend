package user

import (
	"Rabbit-OJ-Backend/controllers/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	userRouter := baseRouter.Group("/user")

	userRouter.GET("/info/:username", Info)
	userRouter.GET("/avatar/:uid", Avatar)
	userRouter.GET("/my", middlewares.AuthJWT(true), My)
	userRouter.GET("/token", middlewares.AuthJWT(true), Token)
	userRouter.POST("/my/avatar", middlewares.AuthJWT(true), UploadAvatar)
	userRouter.POST("/login", Login)
	userRouter.POST("/register", Register)
}
