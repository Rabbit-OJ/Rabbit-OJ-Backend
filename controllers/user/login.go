package user

import "github.com/gin-gonic/gin"

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

func Login(context *gin.Context) {
	var loginForm LoginForm

	if err := context.BindJSON(&loginForm); err != nil {
		context.JSON(500, gin.H{
			"message": "Invalid JSON format",
		})
	}


}
