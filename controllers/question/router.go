package question

import (
	"Rabbit-OJ-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")

	questionRouter.GET("/list/:page", List)

	questionRouter.GET("/item/:tid", Detail)
	questionRouter.POST("/item", middlewares.AuthJWT(), Create)
	questionRouter.PUT("/item/:tid", middlewares.AuthJWT(), Edit)
	questionRouter.DELETE("/item/:tid", middlewares.AuthJWT(), Delete)
}
