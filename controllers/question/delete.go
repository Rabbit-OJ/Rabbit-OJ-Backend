package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	tid, err := strconv.ParseUint(c.Param("tid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if err := QuestionService.Delete(uint32(tid)); err != nil {
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
