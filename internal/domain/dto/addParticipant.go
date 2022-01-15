package dto

type AddParticipant struct {
	ParticipantEmail string `json:"participant_email"`
	RoomID           uint   `json:"room_id"`
}
