package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func MyInfo(c *gin.Context) {
	cid := c.Param("cid")

	contestRaw, ok := c.Get("contest")
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "internal error",
		})

		return
	}

	contest, ok := contestRaw.(*models.Contest)
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "internal error",
		})

		return
	}

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	contestMyInfo, err := ContestService.MyInfo(authObject.Uid, cid, contest)
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
