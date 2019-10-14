package user

import (
	"Rabbit-OJ-Backend/auth"
	UserService "Rabbit-OJ-Backend/services/user"
	"github.com/gin-gonic/gin"
)

func My(c *gin.Context) {
	authObjectRaw, _ := c.Get("AuthObject")
	authObject, ok := authObjectRaw.(*auth.Claims)
	if !ok {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Invalid token",
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
