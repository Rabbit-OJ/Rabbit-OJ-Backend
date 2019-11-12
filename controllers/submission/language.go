package submission

import (
	"Rabbit-OJ-Backend/utils"
	"github.com/gin-gonic/gin"
)

func Language(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": utils.SupportLanguage,
	})
}
