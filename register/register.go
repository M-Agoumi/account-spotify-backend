package register

import (
	"encoding/json"
	"magoumi/spotify-account/config"
	"magoumi/spotify-account/model"
	"net/http"
)

// Register handles user registration.
type Register struct{}

// Register handles the HTTP request for user registration.
func (h *Register) Register(w http.ResponseWriter, r *http.Request) {
	// Load the body
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if email is taken
	var existingUser model.User
	config.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		http.Error(w, "User already registered", http.StatusBadRequest)
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Something went wrong, please try again later", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("User registered successfully"))
}
