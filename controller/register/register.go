package register

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"magoumi/spotify-account/config"
	"magoumi/spotify-account/model"
	"magoumi/spotify-account/util"
	"net/http"
)

// Register handles user registration.
type Register struct{}

// Register handles the HTTP request for user registration.
// @todo add captcha for all endpoints in this file
func (h *Register) Register(w http.ResponseWriter, r *http.Request) {
	// Load the body
	var user model.User
	err := util.DecodeJSONBody(r, &user)
	if err != nil {
		fmt.Println(err)
		util.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check if email is taken
	var existingUser model.User
	config.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		util.JSONError(w, http.StatusBadRequest, "User already registered")
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		util.JSONError(w, http.StatusInternalServerError, "Something went wrong, please try again later")
		return
	}

	util.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *Register) CheckEmail(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := decodeRequestBody(w, r, &user); err != nil {
		util.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the required fields
	if err := validateUser(user); err != nil {
		util.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	email := *user.Email

	existingUser, err := findUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		util.JSONError(w, http.StatusInternalServerError, "Database error")
		return
	}

	if existingUser.ID != 0 {
		util.JSONError(w, http.StatusBadRequest, "User already registered")
		return
	}

	util.JSONResponse(w, http.StatusOK, map[string]string{"message": "Email is available"})
}

func decodeRequestBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		var invalidUnmarshalError *json.InvalidUnmarshalError
		if errors.As(err, &invalidUnmarshalError) {
			util.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		}
		return err
	}
	return nil
}

func findUserByEmail(email string) (model.User, error) {
	var user model.User
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
