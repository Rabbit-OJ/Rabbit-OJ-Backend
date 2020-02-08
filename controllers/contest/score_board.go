package contest

import (
	"Rabbit-OJ-Backend/models/responses"
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ScoreBoard(c *gin.Context) {
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	page, err := strconv.ParseUint(c.Param("page"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	contestInfo, err := contest.Info(uint32(cid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if contestInfo.Status == contest.RoundWaiting {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Contest Not begin",
		})

		return
	}

	scoreBoard, blocked, err := contest.ScoreBoard(contestInfo, uint32(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}
	if blocked {
		c.JSON(200, gin.H{
			"code": 200,
			"message": responses.ScoreBoardResponse{
				List:    make([]*responses.ScoreBoard, 0),
				Count:   contestInfo.Participants,
				Blocked: true,
			},
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"message": responses.ScoreBoardResponse{
			List:    scoreBoard,
			Count:   contestInfo.Participants,
			Blocked: false,
		},
	})
}
