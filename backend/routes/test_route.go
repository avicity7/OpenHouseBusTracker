package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func TestRoute(r *chi.Mux) {
	r.Route("/test", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(controllers.GetTestController))
	})
}
