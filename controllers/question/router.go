package question

import (
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")
	questionRouter.GET("/list/:page", List)
	questionRouter.GET("/content/:tid", Detail)
}
