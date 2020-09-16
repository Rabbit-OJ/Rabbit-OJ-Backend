package submission

import (
	config2 "Rabbit-OJ-Backend/services/judger/config"
	"github.com/gin-gonic/gin"
)

func Language(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": config2.SupportLanguage,
	})
}
