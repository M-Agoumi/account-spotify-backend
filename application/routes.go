package application

import (
	"fmt"
	"github.com/M-Agoumi/account-spotify-backend/controller/login"
	"github.com/M-Agoumi/account-spotify-backend/controller/register"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func loadRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			fmt.Println(fmt.Errorf("error writing response: %w", err))
			return
		}
	})

	router.Route("/register", register.LoadRegisterRoute)
	router.Route("/login", login.LoadLoginRoute)
	return router
}
