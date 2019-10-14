package question

import (
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	tid := c.Param("tid")

	content, err := QuestionService.Content(tid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": content,
		})
	}
}
