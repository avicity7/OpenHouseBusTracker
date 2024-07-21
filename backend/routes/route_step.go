package routes

import (
	"server/controllers"
	// "github.com/jackc/pgx/v5"
	"github.com/go-chi/chi/v5"
)

func RouteStep(r chi.Router) {
	r.Route("/route-step", func(r chi.Router) {
		r.Get("/", controllers.GetAllRouteSteps)
		r.Get("/get-all-stops", controllers.GetAllStopNames)
		r.Get("/{route_name}/{stop_name}", controllers.GetRouteStepHandler)
		r.Post("/create-route-step", controllers.CreateRouteStepHandler)
		r.Put("/update-route-step", controllers.UpdateRouteStepHandler)
		r.Delete("/delete-route-step/{route_name}/{stop_name}", controllers.DeleteRouteStepHandler)
	})
}
