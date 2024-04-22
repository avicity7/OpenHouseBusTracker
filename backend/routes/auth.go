package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Auth(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/getJWT", http.HandlerFunc(controllers.GetJWT))
	})
}
