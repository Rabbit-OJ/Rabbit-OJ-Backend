package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/controllers/websocket"
	"Rabbit-OJ-Backend/models/forms"
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func ClarifyAdd(c *gin.Context) {
	clarifyAddForm := &forms.ContestClarifyAdd{}

	if _, err := auth.GetAuthObjRequireAdmin(c); err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	if err := c.BindJSON(&clarifyAddForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	ccid, err := ContestService.ClarifyCreate(clarifyAddForm.Cid, clarifyAddForm.Message)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	go func() {
		contestBroadCast := &ContestService.HubBroadcast{
			Cid:     clarifyAddForm.Cid,
			Type:    "clarify",
			Message: clarifyAddForm.Message,
		}

		websocket.SocketHub.ContestHub.Broadcast <- contestBroadCast
	}()

	c.JSON(200, gin.H{
		"code":    200,
		"message": ccid,
	})

	return
}
