package models

type Chatroom struct {
	ID              uint `gorm:"primaryKey"`
	sender_id       int
	receiver_id     int
	messages_recent *string
	messages_backup *string
}

type Message struct {
	sender_id   string
	receiver_id string
	message     string
}
