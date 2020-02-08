package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SubmissionOne(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	cid, err := strconv.ParseUint(c.Param("cid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}
	sid, err := strconv.ParseUint(c.Param("sid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	submissionInfo, err := ContestService.SubmissionOne(uint32(sid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if (submissionInfo.Cid != uint32(cid) || submissionInfo.Uid != authObject.Uid) && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Access Denied",
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": submissionInfo,
	})
}
