package contest

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	contestRouter := baseRouter.Group("/contest")

	contestRouter.GET("/list/:page", List)
	contestRouter.GET("/info/:cid", Info)
	contestRouter.GET("/submit/:cid/:id", middlewares.AuthJWT(true), middlewares.CheckContest, Submit)
	contestRouter.GET("/submission/list/:cid", middlewares.TryAuthJWT(), SubmissionList)
	contestRouter.GET("/clarify/:cid", middlewares.TryAuthJWT(), middlewares.CheckContest, Clarify)
	contestRouter.GET("/question/:cid", middlewares.TryAuthJWT(), middlewares.CheckContest, Question)
	contestRouter.POST("/register/:cid/:operation", middlewares.AuthJWT(true), Register)
	contestRouter.GET("/my/info/:cid", middlewares.AuthJWT(true), middlewares.CheckContest, MyInfo)
	contestRouter.GET("/scoreboard/:cid/:page", ScoreBoard)
}
