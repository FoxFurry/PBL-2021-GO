package entity

type Message struct {
	ID         uint   `json:"id"`
	SenderName string `json:"sender_name"`
	SenderID   uint   `json:"sender_id"`
	RoomID     uint   `json:"room_id"`
	Data       string `json:"data"`
	Time       int    `json:"time"`
}
