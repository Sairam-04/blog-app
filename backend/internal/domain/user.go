package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password" validate:"required"`
	Bio            string    `json:"bio" validate:"required"`
	ProfilePicture string    `json:"profile_pic" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id uuid.UUID) (*User, error)
	IsEmailTaken(email string) (bool, error)
	GetUser(name, value string) (*User, error)
}
