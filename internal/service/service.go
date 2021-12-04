package service

import (
	"foxy/internal/domain/dto"
	"foxy/internal/domain/entity"
	"foxy/internal/domain/repository"
	"foxy/internal/http/httperr"
	"foxy/internal/infrastructure/hash"
)

type IService interface {
	RegisterUser(user dto.UserRegister) (uint, error)
	GetUsers() ([]entity.User, error)
	GetUser(id uint) (*entity.User, error)

	GetUsersRooms(userID uint) ([]entity.Room, error)
	CreateRoom(userID uint, newRoom entity.Room) (uint, error)

	Authorize(user dto.UserRegister) (uint, error)

	SendMessage(userID uint, roomID uint, data string) (uint, error)
	GetAllMessages(roomID uint) ([]entity.Message, error)
}

type foxyService struct {
	userRepo            repository.IUser
	roomParticipantRepo repository.IRoomParticipant
	messageRepo         repository.IMessage
}

func NewFoxyService() IService {
	return &foxyService{
		userRepo:            repository.NewUserRepository(),
		roomParticipantRepo: repository.NewRoomParticipantRepository(),
		messageRepo:         repository.NewMessageRepository(),
	}
}

func (s *foxyService) RegisterUser(user dto.UserRegister) (uint, error) {
	var err error
	var userEntity entity.User

	mail, _ := s.userRepo.GetUserByMail(user.Email)

	if mail != nil {
		return 0, httperr.NewErrorAlreadyExists()
	}

	userEntity.FullName = user.FullName
	userEntity.Email = user.Email
	userEntity.PasswordHash, err = hash.HashPassword(user.Password)

	if err != nil {
		return 0, httperr.NewErrorInternal()
	}

	newID, err := s.userRepo.RegisterUser(userEntity)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *foxyService) GetUsers() ([]entity.User, error) {
	users, err := s.userRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *foxyService) GetUser(id uint) (*entity.User, error) {
	user, err := s.userRepo.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *foxyService) Authorize(user dto.UserRegister) (uint, error) {
	userEntity, err := s.userRepo.GetUserByMail(user.Email)
	if err != nil {
		return 0, httperr.NewErrorNotFound()
	}

	ok := hash.CheckPasswordHash(user.Password, userEntity.PasswordHash)
	if !ok {
		return 0, httperr.NewErrorUnauthorized()
	}

	return userEntity.ID, nil
}

func (s *foxyService) GetUsersRooms(userID uint) ([]entity.Room, error) {
	usersRooms, err := s.roomParticipantRepo.GetUsersRooms(userID)

	if err != nil {
		return nil, err
	}

	return usersRooms, nil
}

func (s *foxyService) CreateRoom(userID uint, newRoom entity.Room) (uint, error) {
	newRoomID, err := s.roomParticipantRepo.CreateRoom(userID, newRoom)

	if err != nil {
		return 0, err
	}

	return newRoomID, nil
}

func (s *foxyService) SendMessage(userID uint, roomID uint, data string) (uint, error) {
	newMessageID, err := s.messageRepo.SendMessage(userID, roomID, data)

	if err != nil {
		return 0, err
	}

	return newMessageID, nil
}

func (s *foxyService) GetAllMessages(roomID uint) ([]entity.Message, error) {
	messages, err := s.messageRepo.GetAllMessages(roomID)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *foxyService) GetAllMessagesByUser(roomID uint, userID uint) ([]entity.Message, error) {
	messages, err := s.messageRepo.GetAllMessagesByUser(userID, roomID)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
