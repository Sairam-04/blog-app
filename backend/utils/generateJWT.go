package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	expTime := time.Now().Add(72 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	fmt.Println("secret", jwtSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token", token)

	return token.SignedString([]byte(jwtSecret))
}
