package middlewares

import (
	"Rabbit-OJ-Backend/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func SecretCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		split := strings.Split(token, ",")
		if len(split) != 2 {
			c.JSON(403, gin.H{
				"code":    403,
				"message": "Access Denied",
			})
			c.Abort()
			return
		}

		check, random := split[0], split[1]
		if utils.Md5(random + utils.Secret) != check {
			c.JSON(403, gin.H{
				"code":    403,
				"message": "Access Denied",
			})
			c.Abort()
			return
		}
	}
}