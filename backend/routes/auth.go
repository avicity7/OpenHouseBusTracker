// Package routes defines all the routes for the API and what controllers each route will call.
package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Auth sets up the authentication routes for the given chi router.
func Auth(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		// GET /auth/get-user/{email}
		// Retrieves the user with the specified email.
		r.Get("/get-user/{email}", http.HandlerFunc(controllers.GetUser))

		// GET /auth/get-user-settings/{email}
		// Retrieves the user settings for the user with the specified email.
		r.Get("/get-user-settings/{email}", http.HandlerFunc(controllers.GetUserSettings))

		// GET /auth/verify/{token}
		// Verifies the email using the specified token.
		r.Get("/verify/{token}", http.HandlerFunc(controllers.VerifyEmail))

		// POST /auth/create-user
		// Creates a new user.
		r.Post("/create-user", http.HandlerFunc(controllers.CreateUser))

		// POST /auth/bulk-create-users
		// Creates multiple users in bulk.
		r.Post("/bulk-create-users", http.HandlerFunc(controllers.BulkCreateUsers))

		// POST /auth/login
		// Logs in the user.
		r.Post("/login", http.HandlerFunc(controllers.Login))

		// POST /auth/start-reset-password
		// Starts the password reset process for the user.
		r.Post("/start-reset-password", http.HandlerFunc(controllers.StartResetPassword))

		// POST /auth/reset-password
		// Resets the password for the user.
		r.Post("/reset-password", http.HandlerFunc(controllers.ResetPassword))
	})
}
