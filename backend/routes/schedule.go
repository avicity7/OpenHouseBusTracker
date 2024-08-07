package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Schedules is the function that defines the routes related to creating Bus Schedules.
func Schedules(r *chi.Mux) {
	r.Route("/schedules", func(r chi.Router) {
		r.Get("/get-schedules", http.HandlerFunc(controllers.GetSchedule))
		r.Get("/get-schedule/{id}", http.HandlerFunc(controllers.GetScheduleByID))
		r.Get("/get-dropdown-data", http.HandlerFunc(controllers.GetDropdownData))
		r.Get("/get-user-schedule/{email}", http.HandlerFunc(controllers.GetScheduleByUser))
		r.Get("/get-future-user-schedule/{email}", http.HandlerFunc(controllers.GetFutureScheduleByUser))
		r.Post("/add-schedule", http.HandlerFunc(controllers.CreateBusSchedule))
		r.Put("/update-schedule", http.HandlerFunc(controllers.UpdateBusSchedule))
		r.Put("/update-schedule-route", http.HandlerFunc(controllers.UpdateScheduleRoutes))
		r.Delete("/delete-schedule", http.HandlerFunc(controllers.DeleteBusSchedule))
	})
}
