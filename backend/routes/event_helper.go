package routes

import (
	"net/http"
	"server/controllers"
	"github.com/go-chi/chi/v5"
)

func EventHelper(r chi.Router) {
    r.Route("/event-helpers", func(r chi.Router) {
        r.Get("/get-helpers", http.HandlerFunc(controllers.GetEventHelpers))
		r.Post("/create-helper", http.HandlerFunc(controllers.CreateEventHelper))
		r.Put("/update-helper", http.HandlerFunc(controllers.UpdateEventHelper))
		r.Delete("/delete-helper", http.HandlerFunc(controllers.DeleteEventHelper))
    })
}