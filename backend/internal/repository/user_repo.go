package repository

import (
	"database/sql"

	"github.com/Sairam-04/blog-app/backend/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password, bio, profile_pic, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.Bio, user.ProfilePicture, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) IsEmailTaken(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
