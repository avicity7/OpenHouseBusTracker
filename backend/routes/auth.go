package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Auth(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-user/{email}", http.HandlerFunc(controllers.GetUser))
		r.Get("/get-user-settings/{email}", http.HandlerFunc(controllers.GetUserSettings))
		r.Get("/verify/{token}", http.HandlerFunc(controllers.VerifyEmail))
		r.Post("/create-user", http.HandlerFunc(controllers.CreateUser))
		r.Post("/bulk-create-users", http.HandlerFunc(controllers.BulkCreateUsers))
		r.Post("/login", http.HandlerFunc(controllers.Login))
		r.Post("/start-reset-password", http.HandlerFunc(controllers.StartResetPassword))
		r.Post("/reset-password", http.HandlerFunc(controllers.ResetPassword))
	})
}
