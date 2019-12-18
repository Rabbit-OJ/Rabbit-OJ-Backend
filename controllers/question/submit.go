package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/question"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/utils/files"
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
	questionJudge, err := question.JudgeInfo(tid)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	questionDetail, err := question.Detail(tid)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	if questionDetail.Hide && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Permission Denied",
		})

		return
	}

	fileName, err := files.CodeGenerateFileNameWithMkdir(authObject.Uid)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	filePath, err := files.CodePath(fileName)
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

	submission, err := SubmissionService.Create(tid, authObject.Uid, submitForm.Language, fileName)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	go func() { _ = SubmissionService.Starter([]byte(submitForm.Code), submission, questionJudge, questionDetail) }()
	go question.UpdateAttemptCount(tid)

	c.JSON(200, gin.H{
		"code":    200,
		"message": submission.Sid,
	})
}
