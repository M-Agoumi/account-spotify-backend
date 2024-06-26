package login

import (
	"fmt"
	"github.com/M-Agoumi/account-spotify-backend/model"
	"github.com/M-Agoumi/account-spotify-backend/model/repository"
	"github.com/M-Agoumi/account-spotify-backend/service/jwtService"
	"github.com/M-Agoumi/account-spotify-backend/util"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Login struct{}

// Login
// @todo add captcha for all endpoints in this file
func (h *Login) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := util.DecodeJSONBody(r, &user)
	if err != nil {
		fmt.Println(err)
		util.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	println(user.Password)
	err, isValid := ValidateLoginBody(user)
	if !isValid {
		util.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check if we should query by email or username
	var existingUser model.User
	if model.IsValidEmail(*user.Username) {
		existingUser, _ = repository.FindUserByEmail(*user.Username)
	} else {
		existingUser, _ = repository.FindUserByUsername(*user.Username)
	}

	if err != nil {
		fmt.Println(err)
		util.JSONError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	if existingUser.ID == 0 {
		util.JSONError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// we found the user now let's see if password match
	if checkPasswordHash(user.Password, existingUser.Password) == false {
		util.JSONError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	token, err := jwtService.GenerateJWT(existingUser.ID, *existingUser.Email)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}

	fmt.Printf("Generated JWT: %s\n", token)

	util.JSONResponse(w, http.StatusCreated, map[string]string{"token": token})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
