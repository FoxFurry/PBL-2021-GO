package entity

type Room struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Messages []Message `json:"-" gorm:"foreignKey:RoomID"`
}
