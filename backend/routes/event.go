package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Event(r chi.Router) {
	r.Route("/event", func(r chi.Router) {
		r.Put("/create-follow-bus", http.HandlerFunc(controllers.CreateFollowBus))
		r.Put("/create-event", http.HandlerFunc(controllers.CreateEvent))
		r.Get("/get-follow-bus/{email}", http.HandlerFunc(controllers.GetFollowBus))
		r.Get("/get-all-follow-bus", http.HandlerFunc(controllers.GetAllFollowBus))
		r.Get("/get-buses", http.HandlerFunc(controllers.GetBuses))
		r.Get("/get-current-buses", http.HandlerFunc(controllers.GetCurrentBuses))
		r.Get("/get-route-steps/{routeName}", http.HandlerFunc(controllers.GetRouteSteps))
		r.Get("/get-events/{scheduleId}", http.HandlerFunc(controllers.GetEvents))
		r.Delete("/delete-follow-bus/{email}", http.HandlerFunc(controllers.DeleteFollowBus))
		r.Delete("/delete-last-event/{bus_id}", http.HandlerFunc(controllers.DeleteLastEvent))
	})
}
