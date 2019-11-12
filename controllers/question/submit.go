package question

import (
	"Rabbit-OJ-Backend/auth"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/question"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Submit(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	submitForm := &forms.SubmitForm{}
	if err := c.ShouldBindJSON(submitForm); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	tid := c.Param("tid")
	if _, err := question.Detail(tid); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	fileName, err := utils.CodeGenerateFileNameWithMkdir(authObject.Uid)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	filePath, err := utils.CodePath(fileName)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	if err := ioutil.WriteFile(filePath, []byte(submitForm.Code), 0644); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	sid, err := SubmissionService.Create(tid, authObject.Uid, submitForm.Language, fileName)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	go question.UpdateAttemptCount(tid)
	// todo: add websocket to deliver state
	c.JSON(200, gin.H{
		"code":    200,
		"message": sid,
	})
}
