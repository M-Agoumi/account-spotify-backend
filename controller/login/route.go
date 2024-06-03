package login

import "github.com/go-chi/chi/v5"

func LoadLoginRoute(router chi.Router) {
	registerHandler := Login{}

	router.Post("/", registerHandler.Login)
}
