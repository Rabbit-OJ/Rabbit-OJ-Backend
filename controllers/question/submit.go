package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/controllers/common"
	"Rabbit-OJ-Backend/models/forms"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Submit(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
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

	tid, err := strconv.ParseUint(c.Param("tid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	sid, err := common.CodeSubmit(uint32(tid), submitForm, authObject, false)
	if err != nil {
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
