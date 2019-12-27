package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/responses"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Record(c *gin.Context) {
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

	tid, uid := c.Param("tid"), authObject.Uid

	detail, err := QuestionService.Detail(tid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if detail.Hide && !authObject.IsAdmin {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "Permission Denied",
		})

		return
	}

	list, err := QuestionService.Record(uid, tid, uint32(page))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	count, err := QuestionService.RecordCount(uid, tid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	response := &responses.SubmissionListResponse{
		List:  list,
		Count: count,
	}

	c.JSON(200, gin.H{
		"code": 200,
		"message": response,
	})
}
