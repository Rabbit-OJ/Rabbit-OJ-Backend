package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		writer := context.Writer
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		writer.Header().Add("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,X-Auth-Token")
		writer.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		method := context.Request.Method
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		context.Next()
	}
}
