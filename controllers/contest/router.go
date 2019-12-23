package contest

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	contestRouter := baseRouter.Group("/contest")

	contestRouter.GET("/list/:page", List)
	contestRouter.GET("/info/:cid", Info)
	contestRouter.GET("/question/:cid", middlewares.TryAuthJWT(), Question)
	contestRouter.GET("/register/:cid/:operation", middlewares.AuthJWT(false), Register)
}
