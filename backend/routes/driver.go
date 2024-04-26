package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Driver(r *chi.Mux) {
	r.Route("/driver", func(r chi.Router) {
		r.Get("/get-driver", http.HandlerFunc(controllers.GetDriver))
		r.Post("/add-driver", http.HandlerFunc(controllers.AddDriver))
        r.Put("/update-driver/{id}", http.HandlerFunc(controllers.UpdateDriver))
        r.Delete("/delete-driver/{id}", http.HandlerFunc(controllers.DeleteDriver))
	})
}