package user

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/utils/files"
	"github.com/gin-gonic/gin"
	"os"
)

func Avatar(c *gin.Context) {
	uid := c.Param("uid")

	avatarPath, _ := files.AvatarPath(uid)

	c.Writer.WriteHeader(200)
	c.Header("Content-Disposition", "attachment; filename=avatar.png")
	c.Header("Content-Type", "application/octet-stream")

	if _, err := os.Stat(avatarPath); err != nil {
		defaultPath, _ := files.DefaultAvatarPath()
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

	avatarPath, err := files.AvatarPath(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

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
