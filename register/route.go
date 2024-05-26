package register

import (
	"github.com/go-chi/chi/v5"
)

func LoadRegisterRoute(router chi.Router) {
	registerHandler := Register{}

	router.Post("/", registerHandler.Register)
}
