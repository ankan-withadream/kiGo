package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ankan-withadream/kiGo/internal/config"
	"github.com/ankan-withadream/kiGo/internal/db"
	"github.com/ankan-withadream/kiGo/src/api/models"
	"github.com/ankan-withadream/kiGo/src/services"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var aiclient = &http.Client{}

func Handle_empty(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		c.Writer.Write([]byte("Get Response"))
	case http.MethodPost:
		fmt.Fprint(c.Writer, "POST request")
	default:
		fmt.Fprintf(c.Writer, "Unsupported method: %s", c.Request.Method)
	}
}

func Handle_kigo(c *gin.Context) {
	fmt.Println(c.Request.Method)

	// if GET method, send this
	if c.Request.Method == http.MethodGet {
		c.Writer.Write([]byte("Ei cholche Go :)"))
	}

	// if POST method, do this
	if c.Request.Method == http.MethodPost {

		// read req body as json
		jsondata, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
		}

		// json to payload data
		data := &services.Payload{}
		err = json.Unmarshal(jsondata, data)
		if err != nil {
			fmt.Println("Error parsing json data: ", err)
		}

		// request to AI server and get response
		resp, err := services.Ask_AI(config.APP_CONFIG.AI_CLIENT_API_ENDPOINT, data, aiclient)
		if err != nil {
			fmt.Println("Error from the AI server: ", err)
		}

		// send response to client
		// fmt.Fprint(c.Writer, resp)
		// println(resp)
		c.JSON(http.StatusOK, resp)

	}
}

func Handle_allChats(c *gin.Context) {
	db := db.Get()
	chatrooms := []models.Chatroom{}
	db.Find(&chatrooms)
	c.JSON(http.StatusOK, chatrooms)

}

func Handle_latestChats(c *gin.Context) {
	db := db.Get()
	chatrooms := []models.Chatroom{}
	db.Limit(10).Find(&chatrooms)
	c.JSON(http.StatusOK, chatrooms)

}

func Handle_WS(c *gin.Context) {
	hub := services.New_hub()
	services.Serve_chat_ws(c, hub)
}

func Handle_addNewMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing JSON: %v", err)})
		return
	}

	db := db.Get()
	var chatroom models.Chatroom
	if err := db.First(&chatroom, models.Chatroom{Sender_id: message.Sender_id, Receiver_id: message.Receiver_id}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			chatroom = createNewChatroom(db, &message)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error retrieving chatroom: %v", err)})
			return
		}
	}

	chatroom.Messages_recent = &message.Message

	chatroom.Messages_backup = append(chatroom.Messages_backup, models.MessageBackup{
		IsSender: message.Sender_id == chatroom.Sender_id,
		Message:  message.Message,
	})

	if err := db.Save(&chatroom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving chatroom: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message added to chatroom"})
}

func createNewChatroom(db *gorm.DB, message *models.Message) models.Chatroom {
	if message == nil || message.Message == "" {
		panic("message cannot be nil")
	}

	var sender models.User
	if err := db.Where("UserID = ?", message.Sender_id).First(&sender).Error; err != nil {
		panic(fmt.Sprintf("failed to get sender user: %v", err))
	}

	var receiver models.User
	if err := db.Where("UserID = ?", message.Receiver_id).First(&receiver).Error; err != nil {
		panic(fmt.Sprintf("failed to get receiver user: %v", err))
	}

	chatroom := models.Chatroom{
		Sender_id:       message.Sender_id,
		Sender_name:     &sender.Name,
		Receiver_id:     message.Receiver_id,
		Receiver_name:   &receiver.Name,
		Messages_recent: &message.Message,
		Messages_backup: []models.MessageBackup{{IsSender: true, Message: message.Message}},
	}

	if err := db.Create(&chatroom).Error; err != nil {
		panic(fmt.Sprintf("failed to create chatroom: %v", err))
	}
	return chatroom
}
