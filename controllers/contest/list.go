package contest

import (
	"Rabbit-OJ-Backend/models/responses"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	page, err := strconv.ParseInt(c.Param("page"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	list, err := ContestService.List(int(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	count, err := ContestService.ListCount()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	response := &responses.ContestList{
		List:  list,
		Count: count,
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": response,
	})
}
