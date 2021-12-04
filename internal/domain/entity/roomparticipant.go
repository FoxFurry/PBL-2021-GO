package entity

type Role string

const (
	Admin     Role = "administrator"
	Moderator Role = "moderator"
	Regular   Role = "regular"
)

type RoomParticipant struct {
	ID       uint
	UserID   uint
	RoomID   uint
	UserRole Role
}
