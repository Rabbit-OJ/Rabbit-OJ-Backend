package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func SubmissionList(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	cid := c.Param("cid")
	submissionList, err := ContestService.SubmissionList(authObject.Uid, cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": submissionList,
	})
}
