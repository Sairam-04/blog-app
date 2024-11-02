package utils

import (
	"os"
	"time"

	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(userID string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	expTime := time.Now().Add(72 * time.Hour)

	claims := &types.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
