package question

import (
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	page, err := strconv.ParseUint(c.Param("page"), 10, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	list, err := QuestionService.List(uint32(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"message": list,
	})
}
