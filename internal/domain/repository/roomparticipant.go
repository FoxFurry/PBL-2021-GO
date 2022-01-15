package repository

import (
	"log"

	"foxy/internal/db"
	"foxy/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type IRoomParticipant interface {
	GetUsersRooms(userID uint) ([]entity.Room, error)
	CreateRoom(userID uint, newRoom entity.Room) (uint, error)
	GetRoomParticipants(roomID uint) ([]entity.RoomParticipant, error)
	AddParticipantToRoom(userID, roomID uint) (uint, error)
	DeleteParticipantFromRoom(userID, roomID uint) error
}

type roomParticipantRepository struct{}

func NewRoomParticipantRepository() IRoomParticipant {
	return &roomParticipantRepository{}
}

func (r *roomParticipantRepository) CreateRoom(userID uint, newRoom entity.Room) (uint, error) {
	//CREATE ROOM
	stmt, err := db.GetDB().Prepare(`INSERT INTO room(name) values(?) RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(newRoom.Name).Scan(&newRoom.ID)
	if err != nil {
		return 0, err
	}
	var newRoomParticipant entity.RoomParticipant

	newRoomParticipant.RoomID = newRoom.ID
	newRoomParticipant.UserID = userID
	newRoomParticipant.UserRole = entity.Admin

	//CREATE ROOM PARTICIPANT
	stmt, err = db.GetDB().Prepare(`INSERT INTO room_participant(user_id, room_id, role) values(?,?,?) RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(newRoomParticipant.UserID, newRoomParticipant.RoomID, newRoomParticipant.UserRole).Scan(&newRoomParticipant.ID)
	if err != nil {
		return 0, err
	}

	return newRoomParticipant.RoomID, nil
}

func (r *roomParticipantRepository) GetUsersRooms(userID uint) ([]entity.Room, error) {
	stmt, err := db.GetDB().Prepare(`SELECT room.id, room.name FROM room_participant INNER JOIN room ON room.id = room_id WHERE user_id=$1`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []entity.Room

	for rows.Next() {
		var tempRoom entity.Room

		err = rows.Scan(&tempRoom.ID, &tempRoom.Name)

		if err != nil {
			log.Printf("Unable to scan the room: %v", err)
			continue
		}

		rooms = append(rooms, tempRoom)
	}

	return rooms, nil
}

func (r *roomParticipantRepository) GetRoomParticipants(roomID uint) ([]entity.RoomParticipant, error) {
	stmt, err := db.GetDB().Prepare(`SELECT user.full_name, room_participant.user_id, room_participant.room_id, room_participant.role FROM room_participant INNER JOIN user ON user.id = user_id WHERE room_id=$1`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []entity.RoomParticipant

	for rows.Next() {
		var tempRoom entity.RoomParticipant

		err = rows.Scan(&tempRoom.Name, &tempRoom.UserID, &tempRoom.RoomID, &tempRoom.UserRole)

		if err != nil {
			log.Printf("Unable to scan the room: %v", err)
			continue
		}

		rooms = append(rooms, tempRoom)
	}

	return rooms, nil
}

func (r *roomParticipantRepository) AddParticipantToRoom(userID, roomID uint) (uint, error) {
	stmt, err := db.GetDB().Prepare(`INSERT INTO room_participant(user_id, room_id, role) values(?,?,?) RETURNING id`)
	if err != nil {
		return 0, err
	}

	var newUserID uint
	err = stmt.QueryRow(userID, roomID, entity.Regular).Scan(&newUserID)
	if err != nil {
		return 0, err
	}

	return newUserID, nil
}

func (r *roomParticipantRepository) DeleteParticipantFromRoom(userID, roomID uint) error {
	return nil
}
