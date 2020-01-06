package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/forms"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	questionForm := &forms.QuestionForm{}
	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}
	if err := c.BindJSON(&questionForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	tid := c.Param("tid")
	if err := QuestionService.Edit(tid, questionForm.Subject, questionForm.Content, questionForm.Difficulty, questionForm.TimeLimit, questionForm.SpaceLimit); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 200,
	})
}
