package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Event(r chi.Router) {
	r.Route("/event", func(r chi.Router) {
		r.Put("/create-follow-bus", http.HandlerFunc(controllers.CreateFollowBus))
		r.Get("/get-follow-bus/{email}", http.HandlerFunc(controllers.GetFollowBus))
		r.Delete("/delete-follow-bus/{email}", http.HandlerFunc(controllers.DeleteFollowBus))
		r.Get("/get-buses", http.HandlerFunc(controllers.GetBuses))
		r.Get("/get-route-steps/{routeName}", http.HandlerFunc(controllers.GetRouteSteps))
		r.Get("/get-events/{scheduleId}", http.HandlerFunc(controllers.GetEvents))
	})
}
