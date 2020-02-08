package contest

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	operation := c.Param("operation")
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

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

	contestInfo, err := contest.Info(uint32(cid))
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if contestInfo.Status != contest.RoundWaiting && contestInfo.Status != contest.RoundStarting {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Contest already ended!",
		})

		return
	}

	uid := authObject.Uid
	userIsRegistered, err := contest.IsRegistered(uid, uint32(cid))
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
		if err := contest.Register(uid, uint32(cid)); err != nil {
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
		if err := contest.Unregister(uid, uint32(cid)); err != nil {
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
