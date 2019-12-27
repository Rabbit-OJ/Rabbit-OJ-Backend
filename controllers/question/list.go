package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/responses"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {
	page, err := strconv.ParseUint(c.Param("page"), 10, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
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

	showHide := authObject.IsAdmin
	list, err := QuestionService.List(uint32(page), showHide)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	count, err := QuestionService.ListCount(showHide)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	response := &responses.QuestionListResponse{
		List:  list,
		Count: count,
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": response,
	})
}
