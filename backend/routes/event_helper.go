package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func EventHelper(r chi.Router) {
	r.Route("/event-helpers", func(r chi.Router) {
		r.Get("/get-helpers", http.HandlerFunc(controllers.GetEventHelpers))
		r.Get("/get-event-dropdown", http.HandlerFunc(controllers.GetEventHelperDropdownData))
		r.Post("/create-helpers", http.HandlerFunc(controllers.CreateEventHelpers))
		r.Put("/update-helper", http.HandlerFunc(controllers.UpdateEventHelper))
		r.Delete("/delete-helper", http.HandlerFunc(controllers.DeleteEventHelper))
	})
}
