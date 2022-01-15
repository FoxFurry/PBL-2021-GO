package entity

type Message struct {
	ID       uint   `json:"id"`
	SenderID uint   `json:"sender_id"`
	RoomID   uint   `json:"room_id"`
	Data     string `json:"data"`
	Time     int    `json:"time"`
}
