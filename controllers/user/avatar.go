package user

import (
	"Rabbit-OJ-Backend/auth"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func Avatar(c *gin.Context) {
	uid := c.Param("uid")

	avatarPath, _ := filepath.Abs("./files/avatar/" + uid + ".avatar")

	c.Writer.WriteHeader(200)
	c.Header("Content-Disposition", "attachment; filename=avatar.png")
	c.Header("Content-Type", "application/octet-stream")

	if _, err := os.Stat(avatarPath); err != nil {
		defaultPath, _ := filepath.Abs("./statics/avatar.png")
		c.File(defaultPath)
	} else {
		c.File(avatarPath)
	}
}

func UploadAvatar(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	uid := authObject.Uid
	_, header, err := c.Request.FormFile("avatar")

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	avatarPath, _ := filepath.Abs("./files/avatar/" + uid + ".avatar")
	if err := c.SaveUploadedFile(header, avatarPath); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
	})
}
