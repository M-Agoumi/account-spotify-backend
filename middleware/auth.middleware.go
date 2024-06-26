package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/M-Agoumi/account-spotify-backend/service/jwtService"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

// JSONErrorResponse represents a JSON error response
type JSONErrorResponse struct {
	Error string `json:"error"`
}

// writeJSONError writes an error message as JSON
func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(JSONErrorResponse{Error: message})
	if err != nil {
		return
	}
}

// AuthMiddleware validates the JWT token and extracts the claims
func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	fmt.Printf("Secret Key a: [%s]\n", secretKey) // Debug: Print the secret key

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeJSONError(w, http.StatusUnauthorized, "Authorization header missing")
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			fmt.Println("Token String:", tokenString) // Debug: Print the token string

			if tokenString == authHeader {
				writeJSONError(w, http.StatusUnauthorized, "Malformed token")
				return
			}

			claims := &jwtService.UserClaim{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			fmt.Printf("Claims: %v\n", claims)           // Debug: Print the claims
			fmt.Printf("Token Valid: %v\n", token.Valid) // Debug: Print if the token is valid

			if err != nil || !token.Valid {
				fmt.Printf("Error: %v\n", err) // Debug: Print the error
				writeJSONError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Add claims to the request context
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
