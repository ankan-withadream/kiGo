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

	go read_pump(hub)
	go write_pump(hub)
}

func read_pump(hub *Hub) {

	for {
		for _, conn := range hub.conections {

			if conn == nil {
				continue
			}

			var msgs []*models.Message
			err := conn.ReadJSON(&msgs)
			if err != nil {
				fmt.Println("error reading connection message: ", err)
				continue
			} else {
				fmt.Println("Read connection messages: ", msgs)
			}

			// assuming conn.ReadJSON() will return a slice of models.Message
			hub.messageQueue = append(hub.messageQueue, msgs...)

		}
	}
}

func write_pump(hub *Hub) {
	for {
		for _, message := range hub.messageQueue {
			if message == nil {
				continue
			}

			receiverConn, ok := hub.conections[message.Receiver_id]
			if !ok || receiverConn == nil {
				continue
			}

			err := receiverConn.WriteJSON(message)
			if err != nil {
				fmt.Println("error writing receiver message: ", err)
				continue
			}

		}
	}
}
