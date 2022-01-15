package entity

type Role string

const (
	Admin     Role = "administrator"
	Moderator Role = "moderator"
	Regular   Role = "regular"
)

type RoomParticipant struct {
	ID       uint   `json:"-"`
	Name     string `json:"name"`
	UserID   uint   `json:"user_id"`
	RoomID   uint   `json:"room_id"`
	UserRole Role   `json:"user_role"`
}
