package user

import (
	"Rabbit-OJ-Backend/controllers/auth"
	UserService "Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func My(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	uid := authObject.Uid
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
