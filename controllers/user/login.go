package user

import (
	"Rabbit-OJ-Backend/auth"
	"Rabbit-OJ-Backend/models"
	UserService "Rabbit-OJ-Backend/services/user"
	"Rabbit-OJ-Backend/utils"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	loginForm := models.LoginForm{}

	if err := context.BindJSON(&loginForm); err != nil {
		context.JSON(400, gin.H{
			"code":    400,
			"message": err,
		})
	}

	user, err := UserService.InfoByUsername(loginForm.Username)
	if err != nil {
		context.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	}

	saltPassword := utils.SaltPasswordWithSecret(loginForm.Password)
	if user != nil && user.Password == saltPassword {
		token, err := auth.SignJWT(user)

		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": err,
			})
		} else {
			context.JSON(200, gin.H{
				"code":    200,
				"message": token,
			})
		}
	} else {
		context.JSON(404, gin.H{
			"code":    404,
			"message": "Username or Password wrong.",
		})
	}
}
