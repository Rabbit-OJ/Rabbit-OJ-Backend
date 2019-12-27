package middlewares

import (
	"Rabbit-OJ-Backend/controllers/auth"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func CheckContest(c *gin.Context) {
	cid := c.Param("cid")

	info, err := ContestService.Info(cid)
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
