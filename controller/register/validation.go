package register

import (
	"fmt"
	"github.com/M-Agoumi/account-spotify-backend/model/user"
)

// validateUserEmail checks if the required fields are present and not empty
func validateUserEmail(user user.User) error {
	if user.Email == nil || *user.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
