package repository

import (
	"github.com/M-Agoumi/account-spotify-backend/config"
	"github.com/M-Agoumi/account-spotify-backend/model"
)

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	result := config.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}
