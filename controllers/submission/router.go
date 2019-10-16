package submission

import (
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	submissionRouter := baseRouter.Group("/submission")

	submissionRouter.GET("/list/:uid/:page", List)
	submissionRouter.GET("/detail/:sid", Detail)
	submissionRouter.GET("/code/:sid", Code)
}
