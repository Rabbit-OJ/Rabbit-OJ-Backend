package user

import (
	"Rabbit-OJ-Backend/models"
	UserService "Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func Info(context *gin.Context) {
	username := context.Param("username")

	user, err := UserService.InfoByUsername(username)

	if err != nil {
		context.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	} else {
		otherUser := models.OtherUser{
			Uid:       user.Uid,
			Username:  user.Username,
			IsAdmin:   user.IsAdmin,
			Attempt:   user.Attempt,
			Accept:    user.Accept,
			LoginAt:   user.LoginAt,
			CreatedAt: user.CreatedAt,
		}
		context.JSON(200, gin.H{
			"code":    200,
			"message": otherUser,
		})
	}
}
