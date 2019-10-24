package question

import (
	"Rabbit-OJ-Backend/models/forms"
	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	questionSubmitForm := &forms.QuestionSubmitForm{}

	if err := c.BindJSON(&questionSubmitForm); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}


}
