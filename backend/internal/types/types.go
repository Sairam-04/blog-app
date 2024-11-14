package types

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserIDKey struct{}

type BlogResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Thumbnail   string    `json:"Thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GeneralResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type BlogsResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Thumbnail   string    `json:"Thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllBlogsResp struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Error   string          `json:"error"`
	Blogs   []BlogsResponse `json:"blogs"`
}

type UserBlogResp struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Error   string         `json:"error"`
	Blogs   []BlogResponse `json:"blogs"`
}

type UpdateBlogReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Thumbnail   string `json:"Thumbnail"`
}

type GetBlogResponse struct {
	GeneralResponse
	Blog *BlogsResponse `json:"blog"`
}

type FileUploadResponse struct {
	FilePath string `json:"file_url"`
	GeneralResponse
}
