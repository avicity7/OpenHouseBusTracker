package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Event sets up the routes for handling events.
func Event(r chi.Router) {
	r.Route("/event", func(r chi.Router) {
		// CreateFollowBus creates a new follow bus entry.
		r.Put("/create-follow-bus", http.HandlerFunc(controllers.CreateFollowBus))

		// CreateEvent creates a new event.
		r.Put("/create-event", http.HandlerFunc(controllers.CreateEvent))

		// GetFollowBus retrieves the follow bus entry for the given email.
		r.Get("/get-follow-bus/{email}", http.HandlerFunc(controllers.GetFollowBus))

		// GetAllFollowBus retrieves all follow bus entries.
		r.Get("/get-all-follow-bus", http.HandlerFunc(controllers.GetAllFollowBus))

		// GetBuses retrieves all buses.
		r.Get("/get-buses", http.HandlerFunc(controllers.GetBuses))

		// GetCurrentBuses retrieves the current buses.
		r.Get("/get-current-buses", http.HandlerFunc(controllers.GetCurrentBuses))

		// GetRouteSteps retrieves the route steps for the given route name.
		r.Get("/get-route-steps/{routeName}", http.HandlerFunc(controllers.GetRouteSteps))

		// GetEvents retrieves the events for the given schedule ID.
		r.Get("/get-events/{scheduleId}", http.HandlerFunc(controllers.GetEvents))

		// DeleteFollowBus deletes the follow bus entry for the given email.
		r.Delete("/delete-follow-bus/{email}", http.HandlerFunc(controllers.DeleteFollowBus))

		// DeleteLastEvent deletes the last event for the given bus ID.
		r.Delete("/delete-last-event/{bus_id}", http.HandlerFunc(controllers.DeleteLastEvent))
	})
}
