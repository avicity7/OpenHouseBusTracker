package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func EventHelper(r chi.Router) {
	r.Route("/event-helpers", func(r chi.Router) {
		r.Get("/get-available-swaps/{email}", http.HandlerFunc(controllers.GetAvailableSwaps))
		r.Get("/get-swap-requests/{email}", http.HandlerFunc(controllers.GetSwapRequests))
		r.Get("/get-helpers", http.HandlerFunc(controllers.GetEventHelpers))
		r.Get("/get-event-dropdown", http.HandlerFunc(controllers.GetEventHelperDropdownData))
		r.Post("/create-helpers", http.HandlerFunc(controllers.CreateEventHelpers))
		r.Post("/create-swap-request", http.HandlerFunc(controllers.CreateSwapRequest))
		r.Put("/update-helper", http.HandlerFunc(controllers.UpdateEventHelper))
		r.Put("/accept-swap-request", http.HandlerFunc(controllers.AcceptSwapRequest))
		r.Delete("/delete-helper", http.HandlerFunc(controllers.DeleteEventHelper))
		r.Delete("/delete-swap-request", http.HandlerFunc(controllers.DeleteSwapRequest))
	})
}
