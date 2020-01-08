package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Detail(c *gin.Context) {
	_tid := c.Param("tid")
	tid, err := strconv.ParseUint(_tid, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	detail, err := QuestionService.Detail(uint32(tid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	if detail.Hide && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Permission Denied",
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": detail,
	})

}
