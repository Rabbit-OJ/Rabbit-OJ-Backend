package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func MyInfo(c *gin.Context) {
	cid := c.Param("cid")

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	contestMyInfo, err := contest.MyInfo(authObject.Uid, cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": contestMyInfo,
	})
}
