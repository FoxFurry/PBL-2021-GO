package entity

type User struct {
	ID           uint   `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}
