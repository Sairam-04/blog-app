package types

import "github.com/golang-jwt/jwt"

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
