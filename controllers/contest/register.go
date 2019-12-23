package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	cid, operation := c.Param("cid"), c.Param("operation")

	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if operation != "cancel" && operation != "reg" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Invalid Operation!",
		})

		return
	}

	contestInfo, err := contest.Info(cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if contestInfo.Status != contest.RoundWaiting {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Contest NOT Pending!",
		})

		return
	}

	uid := authObject.Uid
	userIsRegistered, err := contest.IsRegistered(uid, cid)
	if operation == "cancel" && !userIsRegistered {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "NOT Registered",
		})

		return
	}
	if operation == "reg" && userIsRegistered {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Already Registered",
		})

		return
	}

	if operation == "reg" {
		if err := contest.Register(uid, cid); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": err.Error(),
			})

			return
		} else {
			c.JSON(200, gin.H{
				"code": 200,
			})
		}
	}

	if operation == "cancel" {
		if err := contest.Unregister(uid, cid); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": err.Error(),
			})

			return
		} else {
			c.JSON(200, gin.H{
				"code": 200,
			})
		}
	}
}
