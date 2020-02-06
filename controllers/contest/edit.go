package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/forms"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Edit(c *gin.Context) {
	contestForm := &forms.ContestEditForm{}
	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}
	if err := c.BindJSON(&contestForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	cid := c.Param("cid")
	if err := ContestService.Edit(cid, contestForm); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 200,
	})
}

func EditQuestion(c *gin.Context) {
	contestQuestionForm := &forms.ContestQuestionEditForm{}
	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}
	if err := c.BindJSON(&contestQuestionForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	_cid := c.Param("cid")
	cid, err := strconv.ParseUint(_cid, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	// id should in range [0, N - 1] without gaps
	for i := range contestQuestionForm.Data {
		contestQuestionForm.Data[i].Id = i
	}

	if err := ContestService.EditQuestions(uint32(cid), contestQuestionForm.Data); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 200,
	})
}
