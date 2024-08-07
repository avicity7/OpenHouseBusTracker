package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Bus sets up the routes for managing buses.
func Bus(r chi.Router) {
	r.Route("/bus", func(r chi.Router) {
		// CreateBus creates a new bus with the given car plate number.
		r.Put("/create-bus/{carplate}", http.HandlerFunc(controllers.CreateBus))

		// GetBuses retrieves all buses.
		r.Get("/get-buses", http.HandlerFunc(controllers.GetBuses))

		// GetBusStatus retrieves the status of a specific bus.
		r.Get("/get-bus-status/{bus_id}", http.HandlerFunc(controllers.GetBusStatus))

		// UpdateBusVisibility updates the visibility of a specific bus.
		r.Put("/update-bus-visibility/{bus_id}/{hidden}", http.HandlerFunc(controllers.UpdateBusVisibility))

		// UpdateBus updates the details of a specific bus.
		r.Put("/update-bus/{bus_id}/{carplate}", http.HandlerFunc(controllers.UpdateBus))

		// DeleteBus deletes a specific bus.
		r.Delete("/delete-bus/{bus_id}", http.HandlerFunc(controllers.DeleteBus))
	})
}
