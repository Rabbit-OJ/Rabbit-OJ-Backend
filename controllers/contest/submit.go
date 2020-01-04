package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/controllers/common"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/db"
	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	cid, id := c.Param("cid"), c.Param("id")

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

	question, err := ContestService.QuestionOne(cid, id)
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

	tx := db.DB.Begin()
	sid, err := common.CodeSubmit(tx, question.Tid, submitForm, authObject)
	if err != nil {

		tx.Rollback()
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if err := ContestService.Submit(tx,
		sid, cid, authObject.Uid,
		question.Tid,
		ContestService.CalculateTime(contest)); err != nil {

		tx.Rollback()
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	tx.Commit()
	c.JSON(200, gin.H{
		"code":    200,
		"message": sid,
	})
}
