package question

import (
	"Rabbit-OJ-Backend/controllers/validators"
	"Rabbit-OJ-Backend/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("language", validators.Language); err != nil {
			fmt.Println(err)
		}
		if err := v.RegisterValidation("code", validators.Code); err != nil {
			fmt.Println(err)
		}
	}

	questionRouter.GET("/list/:page", List)
	questionRouter.GET("/record/:tid/:page", middlewares.AuthJWT(), Record)
	questionRouter.GET("/item/:tid", Detail)
	questionRouter.POST("/item", middlewares.AuthJWT(), Create)
	questionRouter.PUT("/item/:tid", middlewares.AuthJWT(), Edit)
	questionRouter.DELETE("/item/:tid", middlewares.AuthJWT(), Delete)
	questionRouter.POST("/submit/:tid", middlewares.AuthJWT(), Submit)
}
