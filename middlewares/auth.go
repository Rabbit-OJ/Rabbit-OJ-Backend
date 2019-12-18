package middlewares

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthJWT(tokenInHeader bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		if tokenInHeader {
			token = c.Request.Header.Get("Authorization")
		} else {
			token = c.PostForm("token")
		}

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

func TryAuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		guestJWT := &auth.Claims{
			Uid:      "",
			Username: "",
			IsAdmin:  false,
		}

		token = strings.ReplaceAll(token, "Bearer ", "")
		if token == "" {
			c.Set("AuthObject", guestJWT)
			return
		}

		claims, err := auth.VerifyJWT(token)
		if err != nil {
			c.Set("AuthObject", guestJWT)
			return
		}

		c.Set("AuthObject", claims)
	}
}
