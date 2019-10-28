package user

import (
	"Rabbit-OJ-Backend/auth"
	"Rabbit-OJ-Backend/models/responses"
	UserService "Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	user, err := UserService.InfoByUid(authObject.Uid)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	token, err := auth.SignJWT(user)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"message": &responses.LoginResponse{
				Token:    token,
				Uid:      user.Uid,
				Username: user.Username,
				IsAdmin:  user.IsAdmin,
			},
		})
	}

}
