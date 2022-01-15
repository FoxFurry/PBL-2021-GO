package repository

import (
	"log"
	"time"

	"foxy/internal/db"
	"foxy/internal/domain/entity"
)

type IMessage interface {
	SendMessage(userID uint, roomID uint, data string) (uint, error)
	GetAllMessages(roomID uint) ([]entity.Message, error)
	GetAllMessagesByUser(userID, roomID uint) ([]entity.Message, error)
}

type messageRepository struct{}

func NewMessageRepository() IMessage {
	return &messageRepository{}
}

func (r *messageRepository) SendMessage(userID uint, roomID uint, data string) (uint, error) {
	stmt, err := db.GetDB().Prepare(`INSERT INTO message(sender_id, room_id, data, time) values(?,?,?,?) RETURNING id`)

	if err != nil {
		return 0, err
	}

	var newMessageID uint
	err = stmt.QueryRow(userID, roomID, data, time.Now().Unix()).Scan(&newMessageID)
	if err != nil {
		return 0, err
	}

	return newMessageID, nil
}

func (r *messageRepository) GetAllMessages(roomID uint) ([]entity.Message, error) {
	stmt, err := db.GetDB().Prepare(`SELECT user.full_name, message.sender_id, message.room_id, message.data, message.time FROM message INNER JOIN user ON user.id = sender_id WHERE room_id=$1`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message

	for rows.Next() {
		var tempMessage entity.Message

		err = rows.Scan(&tempMessage.SenderName, &tempMessage.SenderID, &tempMessage.RoomID, &tempMessage.Data, &tempMessage.Time)
		if err != nil {
			log.Printf("Unable to scan the message: %v", err)
			continue
		}

		messages = append(messages, tempMessage)
	}

	return messages, nil
}

func (r *messageRepository) GetAllMessagesByUser(userID, roomID uint) ([]entity.Message, error) {
	stmt, err := db.GetDB().Prepare(`SELECT *FROM message WHERE room_id=$1 AND user_id=$2`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(roomID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message

	for rows.Next() {
		var tempMessage entity.Message

		err = rows.Scan(&tempMessage.ID, &tempMessage.SenderID, &tempMessage.RoomID, &tempMessage.Data, &tempMessage.Time)

		if err != nil {
			log.Printf("Unable to scan the message: %v", err)
			continue
		}

		messages = append(messages, tempMessage)
	}

	return messages, nil
}
