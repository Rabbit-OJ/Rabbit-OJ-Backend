package user

import (
	UserService "Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func My(c *gin.Context) {
	uid := "1"

	user, err := UserService.MyInfoByUid(uid)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": user,
		})
	}
}
