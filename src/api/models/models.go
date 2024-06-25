package models

type Chatroom struct {
	ID              uint `gorm:"primaryKey"`
	sender_id       int
	sender_name 	*string
	receiver_id     int
	receiver_name 	*string
	messages_recent *string
	messages_backup *string
}

type Message struct {
	Sender_id   string
	Receiver_id string
	Message     string
}
