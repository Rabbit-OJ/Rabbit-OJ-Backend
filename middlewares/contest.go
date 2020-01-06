package middlewares

import (
	"Rabbit-OJ-Backend/controllers/auth"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckContest(c *gin.Context) {
	_cid := c.Param("cid")
	cid, err := strconv.ParseUint(_cid, 10, 32)

	info, err := ContestService.Info(uint32(cid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	if info.Status == 0 && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Permission Denied.",
		})

		c.Abort()
		return
	}

	c.Set("contest", info)
}
