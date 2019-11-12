package question

import (
	"Rabbit-OJ-Backend/middlewares"
	"Rabbit-OJ-Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func languageValidator(
	_ *validator.Validate, _ reflect.Value, _ reflect.Value,
	field reflect.Value,
	_ reflect.Type, _ reflect.Kind, _ string,
) bool {
	if language, ok := field.Interface().(string); ok {
		supportLanguage := false

		for _, lang := range utils.SupportLanguage {
			if language == lang.Value {
				supportLanguage = true
				break
			}
		}
		return supportLanguage
	}
	return false
}

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("language", languageValidator); err != nil {
			fmt.Println(err.Error())
		}
	}

	questionRouter.GET("/list/:page", List)
	questionRouter.GET("/item/:tid", Detail)
	questionRouter.POST("/item", middlewares.AuthJWT(), Create)
	questionRouter.PUT("/item/:tid", middlewares.AuthJWT(), Edit)
	questionRouter.DELETE("/item/:tid", middlewares.AuthJWT(), Delete)
	questionRouter.POST("/submit/:tid", middlewares.AuthJWT(), Submit)

}
