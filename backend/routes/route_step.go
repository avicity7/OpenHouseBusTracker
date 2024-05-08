package routes

import (
	"server/controllers"
	// "github.com/jackc/pgx/v5"
	"github.com/go-chi/chi/v5"
)

func RouteSteps(r chi.Router) {
	r.Route("/route-steps", func(r chi.Router) {
		r.Get("/{route_name}", controllers.GetAllRouteStepsHandler)
		r.Get("/{route_name}/{stop_name}", controllers.GetRouteStepHandler)
		r.Put("/create-route-step", controllers.CreateRouteStepHandler)
		r.Put("/update-route-step", controllers.UpdateRouteStepHandler)
		r.Delete("/{route_name}/{stop_name}", controllers.DeleteRouteStepHandler)
	})
}
