package middlewares

import (
	"Rabbit-OJ-Backend/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		token = strings.ReplaceAll(token, "Bearer ", "")
		if token == "" {
			c.JSON(403, gin.H{
				"code":    403,
				"message": "Access Denied",
			})
			c.Abort()
			return
		}

		claims, err := auth.VerifyJWT(token)

		if err != nil {
			c.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("AuthObject", claims)
	}
}
