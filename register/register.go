package register

import (
	"net/http"
)

// Register handles user registration.
type Register struct{}

// Register handles the HTTP request for user registration.
func (h *Register) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("User registered successfully"))
}
