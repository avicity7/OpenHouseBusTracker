package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// TestRoute has been defined in order to provide a way to do a simple health check.
func TestRoute(r *chi.Mux) {
	r.Route("/test", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(controllers.GetTestController))
	})
}
