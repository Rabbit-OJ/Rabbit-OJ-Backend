package submission

import (
	"Rabbit-OJ-Backend/controllers/auth"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/utils/files"
	"github.com/gin-gonic/gin"
	"os"
)

func Code(c *gin.Context) {
	sid := c.Param("sid")

	authObject, err := auth.GetAuthObjRequireAdmin(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	submission, err := SubmissionService.Detail(sid)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	if authObject.Uid != submission.Uid && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Permission Denied",
		})

		return
	}

	codePath, _ := files.CodePath(submission.FileName)
	c.Writer.WriteHeader(200)
	c.Header("Content-Disposition", "attachment; filename=code.txt")
	c.Header("Content-Type", "text/plain")

	if _, err := os.Stat(codePath); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	} else {
		c.File(codePath)
	}
}
