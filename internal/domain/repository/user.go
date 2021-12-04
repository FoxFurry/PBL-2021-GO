package repository

import (
	"foxy/internal/db"
	"foxy/internal/domain/entity"
	"log"
)

type IUser interface {
	RegisterUser(user entity.User) (uint, error)
	GetUser(id uint) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByMail(mail string) (*entity.User, error)
}

type userRepository struct{}

func NewUserRepository() IUser {
	return &userRepository{}
}

func (r *userRepository) RegisterUser(user entity.User) (uint, error) {
	stmt, err := db.GetDB().Prepare(`INSERT INTO user(full_name, email, password) values(?,?,?) RETURNING id`)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(user.FullName, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *userRepository) GetUser(id uint) (*entity.User, error) {
	user := new(entity.User)

	stmt, err := db.GetDB().Prepare(`SELECT * FROM user WHERE id=$1`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.FullName, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User

	stmt, err := db.GetDB().Prepare(`SELECT * FROM user`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tempUser entity.User
		err = rows.Scan(&tempUser.ID, &tempUser.FullName, &tempUser.Email, &tempUser.PasswordHash)

		if err != nil {
			log.Printf("Unable to scan user from row: %v", err)
			continue
		}

		users = append(users, tempUser)
	}

	return users, nil
}

func (r *userRepository) GetUserByMail(mail string) (*entity.User, error) {
	user := new(entity.User)

	stmt, err := db.GetDB().Prepare(`SELECT * FROM user WHERE email=$1`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(mail).Scan(&user.ID, &user.FullName, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
