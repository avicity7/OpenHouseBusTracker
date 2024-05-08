package routes

import (
	// "github.com/jackc/pgx/v5"
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Route("/routes", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(controllers.GetAllRoutesHandler))
		r.Get("/{name}", http.HandlerFunc(controllers.GetRouteByNameHandler))
	})
}
