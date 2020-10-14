package submission

import (
	"Rabbit-OJ-Backend/controllers/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	submissionRouter := baseRouter.Group("/submission")

	submissionRouter.GET("/list/:uid/:page", List)
	submissionRouter.GET("/detail/:sid", Detail)
	submissionRouter.POST("/code/:sid", middlewares.AuthJWT(false), Code)
	submissionRouter.GET("/language", Language)
}
