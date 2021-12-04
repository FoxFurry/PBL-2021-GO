package dto

type MessageCreate struct {
	RoomID uint   `json:"room_id"`
	Data   string `json:"data"`
}
