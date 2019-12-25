package contest

import (
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ScoreBoard(c *gin.Context) {
	cid := c.Param("cid")

	page, err := strconv.ParseUint(c.Param("page"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	contestInfo, err := contest.Info(cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if contestInfo.Status == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Contest Not begin",
		})

		return
	}

	scoreBoard, err := contest.ScoreBoard(cid, uint32(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": scoreBoard,
	})
}
