package routes

import (
	"net/http"
	"server/controllers"
	"github.com/go-chi/chi/v5"
)

func Driver(r chi.Router) {
	r.Route("/driver", func(r chi.Router) {
		r.Get("/get-driver", http.HandlerFunc(controllers.GetDriver))
		r.Get("/get-driver-hours", http.HandlerFunc(controllers.GetScheduleTimeDiff))
		r.Post("/add-driver", http.HandlerFunc(controllers.AddDriver))
		r.Put("/update-driver/{id}", http.HandlerFunc(controllers.UpdateDriver))
		r.Delete("/delete-driver/{id}", http.HandlerFunc(controllers.DeleteDriver))
	})
}
