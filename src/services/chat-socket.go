package services

import (
	"fmt"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin: func(r *http.Request) bool {
	// 	return true
	// },
}

func Serve_chat_ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Request.Response.Header)
	if err != nil {
		fmt.Println("socket conection upgrade error:", err)
		return
	}
	defer conn.Close()  
	
}