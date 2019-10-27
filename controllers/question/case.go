package question

import (
	"Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

func Case(c *gin.Context) {
	tid := c.Param("tid")

	testCase, err := question.Case(tid)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	body, err := proto.Marshal(testCase)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.ProtoBuf(200, body)
}
