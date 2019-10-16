package submission

import (
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func Code(c *gin.Context) {
	sid := c.Param("sid")

	submission, err := SubmissionService.Detail(sid)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	avatarPath, _ := filepath.Abs("./files/submission/" + submission.FileName)
	c.Writer.WriteHeader(200)
	c.Header("Content-Disposition", "attachment; filename=code.txt")
	c.Header("Content-Type", "text/plain")

	if _, err := os.Stat(avatarPath); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	} else {
		c.File(avatarPath)
	}
}
