package user

import (
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerForm := forms.RegisterForm{}

	if err := c.BindJSON(&registerForm); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})
	}

	uid, err := user.Register(&registerForm)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": uid,
		})
	}
}
