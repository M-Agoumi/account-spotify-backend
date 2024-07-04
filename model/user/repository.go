package user

import (
	"github.com/M-Agoumi/account-spotify-backend/config"
)

func FindUserByEmail(email string) (*User, error) {
	var user User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
