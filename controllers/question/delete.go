package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	tid := c.Param("tid")
	if err := QuestionService.Delete(tid); err != nil {
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
