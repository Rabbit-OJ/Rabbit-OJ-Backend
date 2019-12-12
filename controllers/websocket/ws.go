package websocket

import (
	"github.com/gin-gonic/gin"
)

var (
	SocketHub *Hub
)

func WebSocket(baseRouter *gin.Engine) {
	SocketHub = newHub()

	go SocketHub.run()
	baseRouter.GET("/ws/:sid", serveWs)
}
