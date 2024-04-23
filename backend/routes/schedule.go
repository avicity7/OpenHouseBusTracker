package routes

import (
    "net/http"
    "server/controllers"

    "github.com/go-chi/chi/v5"
)

func Schedules(r *chi.Mux) {
    r.Route("/schedules", func(r chi.Router) {
        r.Get("/get-schedule", http.HandlerFunc(controllers.GetSchedule))
    })
}
