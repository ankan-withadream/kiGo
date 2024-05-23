package services

import (
	"fmt"

	"github.com/ankan-withadream/kiGo/src/api/models"
	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

type Hub struct {
	conections   map[string]*websocket.Conn
	messageQueue []*models.Message
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin: func(r *http.Request) bool {
	// 	return true
	// },
}

func New_hub() *Hub {
	return &Hub{
		conections:   make(map[string]*websocket.Conn),
		messageQueue: make([]*models.Message, 100),
	}
}

func Serve_chat_ws(c *gin.Context, hub *Hub) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, c.Request.Response.Header)
	if err != nil {
		fmt.Println("socket conection upgrade error:", err)
		return
	}
	defer conn.Close()

}

func read_pump(hub *Hub) {

	for {
		for _, conn := range hub.conections {

			_, messages, err = conn.ReadMessage()
			if err != nil {
				fmt.Println("error reading connection message: ", err)
			}
			
			for _, message := range messages {
				append(hub.messageQueue, message)
			}

		}
	}
}

func write_pump(hub *Hub) {
	for {
		for _, message := range hub.messageQueue {

		}
	}
}
