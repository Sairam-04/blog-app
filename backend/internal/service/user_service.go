package service

import (
	"errors"
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
	"github.com/Sairam-04/blog-app/backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(user *domain.User) (string, error) {

	exists, err := s.userRepo.IsEmailTaken(user.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("email already registered")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}
	return token, nil

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
