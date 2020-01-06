package contest

import (
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Question(c *gin.Context) {
	_cid := c.Param("cid")
	cid, err := strconv.ParseUint(_cid, 32, 10)
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
