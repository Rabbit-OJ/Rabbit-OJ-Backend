package contest

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	contestRouter := baseRouter.Group("/contest")

	contestRouter.GET("/list/:page", List)
	contestRouter.GET("/info/:cid", Info)
	contestRouter.POST("/submit/:cid/:tid", middlewares.AuthJWT(true), middlewares.CheckContest, Submit)
	contestRouter.GET("/submission/list/:cid", middlewares.AuthJWT(true), SubmissionList)
	contestRouter.GET("/submission/one/:cid/:sid", middlewares.AuthJWT(true), SubmissionOne)
	contestRouter.GET("/clarify/:cid", middlewares.TryAuthJWT(), middlewares.CheckContest, Clarify)
	contestRouter.POST("/clarify/add", middlewares.AuthJWT(true), ClarifyAdd)
	contestRouter.GET("/question/:cid", middlewares.TryAuthJWT(), middlewares.CheckContest, Question)
	contestRouter.POST("/register/:cid/:operation", middlewares.AuthJWT(true), middlewares.CheckContest, Register)
	contestRouter.GET("/my/info/:cid", middlewares.AuthJWT(true), middlewares.CheckContest, MyInfo)
	contestRouter.GET("/scoreboard/:cid/:page", ScoreBoard)
	contestRouter.PUT("/info/:cid", middlewares.AuthJWT(true), Edit)
}
