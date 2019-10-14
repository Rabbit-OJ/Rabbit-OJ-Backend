package submission

import (
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	page, err := strconv.ParseUint(c.Param("page"), 10, 32)
	uid := c.Param("uid")

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	list, err := SubmissionService.List(uid, uint32(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"list": list,
	})
}
