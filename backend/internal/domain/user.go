package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profile_pic"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	IsEmailTaken(email string) (bool, error)
	GetUser(name, value string) (*User, error)
}
