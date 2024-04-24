package routes

import (
    "net/http"
    "server/controllers"

    "github.com/go-chi/chi/v5"
)

func Schedules(r *chi.Mux) {
    r.Route("/schedules", func(r chi.Router) {
        r.Get("/get-schedule", http.HandlerFunc(controllers.GetSchedule))
        r.Post("/create-schedule", http.HandlerFunc(controllers.CreateBusSchedule))
        r.Put("/update-schedule", http.HandlerFunc(controllers.UpdateBusSchedule))
        r.Delete("/delete-schedule/{id}", http.HandlerFunc(controllers.DeleteBusSchedule))
    })
}
