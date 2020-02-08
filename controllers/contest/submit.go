package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/controllers/common"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Submit(c *gin.Context) {
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}
	tid, err := strconv.ParseUint(c.Param("tid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

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

	question, err := ContestService.QuestionOne(uint32(cid), uint32(tid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if contest.Status != 1 {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Contest Not start.",
		})

		return
	}

	submitForm := &forms.SubmitForm{}
	if err := c.ShouldBindJSON(submitForm); err != nil {
		c.JSON(404, gin.H{
			"code":    404,
			"message": err.Error(),
		})

		return
	}

	sid, err := common.CodeSubmit(question.Tid, submitForm, authObject, true)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if err := ContestService.Submit(
		sid, uint32(cid), authObject.Uid,
		question.Tid,
		ContestService.CalculateTime(contest)); err != nil {

		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": sid,
	})
}
