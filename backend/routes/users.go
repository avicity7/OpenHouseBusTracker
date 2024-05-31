package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Users(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/get-users", http.HandlerFunc(controllers.GetUsers))
		r.Get("/get-roles", http.HandlerFunc(controllers.GetUserRoles))
		r.Put("/update-user", http.HandlerFunc(controllers.UpdateUserRole))
		r.Delete("/delete-user/{email}", http.HandlerFunc(controllers.DeleteUser))
	})
}
