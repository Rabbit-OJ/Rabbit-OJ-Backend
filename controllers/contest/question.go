package contest

import (
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Question(c *gin.Context) {
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	questions, err := ContestService.QuestionExtended(uint32(cid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": questions,
	})
}
