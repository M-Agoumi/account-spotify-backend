package jwtService

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type UserClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	issuer := os.Getenv("JWT_ISSUER")
	expirationTimeStr := os.Getenv("JWT_EXPIRATION_TIME")

	// Parse expirationTimeStr to time.Duration
	expirationTime, err := time.ParseDuration(expirationTimeStr)
	if err != nil {
		return "", fmt.Errorf("invalid expiration time format: %v", err)
	}
	// Create the custom claims
	claims := UserClaim{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	// Create the token using the HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
