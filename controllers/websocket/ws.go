package websocket

import (
	"github.com/gin-gonic/gin"
)

var (
	SocketHub *Hub
)

func WebSocket(baseRouter *gin.Engine) {
	SocketHub = newHub()

	go SocketHub.JudgeHub.run()
	go SocketHub.ContestHub.run()

	baseRouter.GET("/ws/:sid", serveJudgeWs)
	baseRouter.GET("/contest/ws/:cid/:uid", serveContestWs)
}
