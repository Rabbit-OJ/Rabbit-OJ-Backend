package user

import (
	"Rabbit-OJ-Backend/auth"
	"Rabbit-OJ-Backend/models/forms"
	UserService "Rabbit-OJ-Backend/services/user"
	"Rabbit-OJ-Backend/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginForm := forms.LoginForm{}

	if err := c.BindJSON(&loginForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	user, err := UserService.InfoByUsername(loginForm.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	saltPassword := utils.SaltPasswordWithSecret(loginForm.Password)
	if user != nil && user.Password == saltPassword {
		token, err := auth.SignJWT(user)

		if err != nil {
			c.JSON(500, gin.H{
				"code":    500,
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"message": token,
			})
		}
	} else {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Username or Password wrong.",
		})
	}
}
