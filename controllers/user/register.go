package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	registerForm := models.RegisterForm{}

	if err := context.BindJSON(&registerForm); err != nil {
		context.JSON(404, gin.H{
			"code":    404,
			"message": err,
		})
	}

	uid, err := user.Register(&registerForm)
	if err != nil {
		context.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	} else {
		context.JSON(200, gin.H{
			"code":    200,
			"message": uid,
		})
	}
}
