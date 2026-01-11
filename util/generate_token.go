package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenJWT(id int, email string) (string, error) {
	jwtClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"email": email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := jwtClaim.SignedString([]byte(secretKey))
	
	if err != nil {
		return "", fmt.Errorf("signed string %s", err)
	}

	return tokenString, nil
}