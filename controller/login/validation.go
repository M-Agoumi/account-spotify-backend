package login

import (
	"github.com/M-Agoumi/account-spotify-backend/model/user"
	"github.com/M-Agoumi/account-spotify-backend/util"
)

func ValidateLoginBody(user user.User) (error, bool) {
	validationErrors := util.NewValidationError()

	// Validate username
	username := ""
	if user.Username != nil {
		username = *user.Username
	}

	if username == "" {
		validationErrors.Add("username is mandatory")
	}

	// Validate password
	if user.Password == "" {
		validationErrors.Add("password is mandatory")
	}

	if validationErrors.HasErrors() {
		return validationErrors, false
	}

	return nil, true
}
