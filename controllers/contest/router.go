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
	contestRouter.POST("/register/:cid/:operation", middlewares.AuthJWT(true), Register)
	contestRouter.GET("/my/info/:cid", middlewares.AuthJWT(true), MyInfo)
	contestRouter.GET("/scoreboard/:cid/:page", ScoreBoard)
}
