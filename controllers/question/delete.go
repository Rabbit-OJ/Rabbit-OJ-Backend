package question

import (
	QuestionService "Rabbit-OJ-Backend/services/question"
	"Rabbit-OJ-Backend/utils"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	if _, err := utils.GetAuthObjRequireAdmin(c); err != nil {
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
