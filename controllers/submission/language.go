package submission

import (
	config2 "github.com/Rabbit-OJ/Rabbit-OJ-Judger/config"
	"github.com/gin-gonic/gin"
)

func Language(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": config2.SupportLanguage,
	})
}
