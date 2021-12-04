package dto

type GetAllMessageByUser struct {
	UserID uint `json:"user_id"`
	RoomID uint `json:"room_id"`
}
