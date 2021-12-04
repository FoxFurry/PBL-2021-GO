package entity

type Room struct {
	ID       uint
	Name     string
	Messages []Message `json:"-" gorm:"foreignKey:RoomID"`
}
