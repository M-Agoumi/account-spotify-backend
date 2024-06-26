package account

import (
	"github.com/M-Agoumi/account-spotify-backend/middleware"
	"github.com/go-chi/chi/v5"
	"os"
)

func LoadAccountRoute(router chi.Router) {
	accountHandler := Account{}

	router.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET_KEY")))
	router.Post("/", accountHandler.User)
}
