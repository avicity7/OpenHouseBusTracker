package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Bus(r chi.Router) {
	r.Route("/bus", func(r chi.Router) {
		r.Put("/create-bus/{carplate}", http.HandlerFunc(controllers.CreateBus))
		r.Get("/get-buses", http.HandlerFunc(controllers.GetBuses))
		r.Get("/get-bus-status/{carplate}", http.HandlerFunc(controllers.GetBusStatus))
		r.Put("/update-bus-visibility/{carplate}/{hidden}", http.HandlerFunc(controllers.UpdateBusVisibility))
		r.Delete("/delete-bus/{carplate}", http.HandlerFunc(controllers.DeleteBus))
	})
}
