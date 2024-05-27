package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Auth(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-user/{email}", http.HandlerFunc(controllers.GetUser))
		r.Get("/verify/{token}", http.HandlerFunc(controllers.VerifyEmail))
		r.Post("/create-user", http.HandlerFunc(controllers.CreateUser))
		r.Post("/bulk-create-users", http.HandlerFunc(controllers.BulkCreateUsers))
		r.Post("/login", http.HandlerFunc(controllers.Login))
		r.Post("/verify-refresh/{refresh}", http.HandlerFunc(controllers.VerifyRefresh))
	})
}
