package application

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"magoumi/spotify-account/register"
	"net/http"
)

func loadRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
	return router
}
