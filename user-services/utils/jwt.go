package utils

import (
	"prueba_tecnica_golang/user-services/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	// Replace "your_secret_key" with your actual key
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	return tokenString, err
}
