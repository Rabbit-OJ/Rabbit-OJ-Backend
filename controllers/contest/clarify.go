package contest

import (
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func Clarify(c *gin.Context) {
	cid := c.Param("cid")

	clarify, err := contest.ClarifyList(cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": clarify,
	})
}
