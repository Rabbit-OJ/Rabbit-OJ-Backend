package websocket

import (
	"Rabbit-OJ-Backend/services/contest"
	"Rabbit-OJ-Backend/services/submission"
	"github.com/gin-gonic/gin"
)

var (
	SocketHub *Hub
)

func WebSocket(baseRouter *gin.Engine) {
	SocketHub = newHub()

	go SocketHub.JudgeHub.Run()
	go SocketHub.ContestHub.Run()

	baseRouter.GET("/ws/:sid", submission.ServeJudgeWs(SocketHub.JudgeHub))
	baseRouter.GET("/contest/ws/:cid/:uid", contest.ServeContestWs(SocketHub.ContestHub))
}
