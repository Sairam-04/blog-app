package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Sairam-04/blog-app/backend/internal/types"
	"github.com/Sairam-04/blog-app/backend/utils"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("x-auth-token")
		err := godotenv.Load(".env")
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization Header Missing")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			utils.RespondWithError(w, http.StatusUnauthorized, "malformed token")
			return
		}
		type userIDKey struct{} // Define a custom type for the context key

		var claimsContextKey userIDKey
		claims := &types.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), claimsContextKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
