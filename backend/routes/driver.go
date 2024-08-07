package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Driver sets up the routes for the driver related endpoints.
func Driver(r chi.Router) {
	r.Route("/driver", func(r chi.Router) {
		// GetDrivers returns the list of all drivers.
		r.Get("/get-drivers", http.HandlerFunc(controllers.GetDrivers))

		// GetScheduleTimeDiff returns the time difference between the current time and the scheduled time for a driver.
		r.Get("/get-driver-hours", http.HandlerFunc(controllers.GetScheduleTimeDiff))

		// AddDriver adds a new driver.
		r.Post("/add-driver", http.HandlerFunc(controllers.AddDriver))

		// UpdateDriver updates an existing driver.
		r.Put("/update-driver", http.HandlerFunc(controllers.UpdateDriver))

		// DeleteDriver deletes a driver.
		r.Delete("/delete-driver", http.HandlerFunc(controllers.DeleteDriver))
	})
}
