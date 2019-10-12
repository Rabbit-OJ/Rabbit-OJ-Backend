package user

import (
	. "Rabbit-OJ-Backend/services/user"
	"Rabbit-OJ-Backend/utils"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(context *gin.Context) {
	var loginForm LoginForm

	if err := context.BindJSON(&loginForm); err != nil {
		context.JSON(400, gin.H{
			"code":    400,
			"message": "Invalid format",
		})
	}

	user := InfoByUsername(loginForm.Username)
	inputPassword := utils.SaltPasswordWithSecret(loginForm.Password)

	if user.Password == inputPassword {
		context.JSON(200, gin.H{
			"code":    200,
			"message": "Login Success.",
		})
	} else {
		context.JSON(404, gin.H{
			"code":    404,
			"message": "Username or Password wrong.",
		})
	}
}
