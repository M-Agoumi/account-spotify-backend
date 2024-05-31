package register

import (
	"fmt"
	"magoumi/spotify-account/model"
)

// validateUser checks if the required fields are present and not empty
func validateUser(user model.User) error {
	if user.Email == nil || *user.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
