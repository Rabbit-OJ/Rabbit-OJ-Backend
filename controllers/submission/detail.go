package submission

import (
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Detail(c *gin.Context) {
	_sid := c.Param("sid")
	sid, err := strconv.ParseUint(_sid, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	submission, err := SubmissionService.Detail(uint32(sid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": submission,
		})
	}
}
