package question

import (
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")


	questionRouter.GET("/list/:page", List)

	questionRouter.POST("/item", Detail)
	questionRouter.GET("/item/:tid", Detail)
	questionRouter.PUT("/item/:tid", Detail)
	questionRouter.DELETE("/item/:tid", Detail)
}
