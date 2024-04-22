package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Auth(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-user/{email}", http.HandlerFunc(controllers.GetUser))
		r.Post("/create-user", http.HandlerFunc(controllers.CreateUser))
		r.Post("/login", http.HandlerFunc(controllers.Login))
	})
}
