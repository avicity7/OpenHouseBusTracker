package routes

import (
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// RouteStep is a function that defines the routes for handling route steps.
// It sets up the endpoints for getting all route steps, getting all stop names,
// getting a specific route step, creating a route step, updating a route step,
// and deleting a route step.
func RouteStep(r chi.Router) {
	r.Route("/route-step", func(r chi.Router) {
		// GET /route-step
		// Handler: controllers.GetAllRouteSteps
		r.Get("/", controllers.GetAllRouteSteps)

		// GET /route-step/get-all-stops
		// Handler: controllers.GetAllStopNames
		r.Get("/get-all-stops", controllers.GetAllStopNames)

		// GET /route-step/{route_name}/{stop_name}
		// Handler: controllers.GetRouteStepHandler
		r.Get("/{route_name}/{stop_name}", controllers.GetRouteStepHandler)

		// POST /route-step/create-route-step
		// Handler: controllers.CreateRouteStepHandler
		r.Post("/create-route-step", controllers.CreateRouteStepHandler)

		// PUT /route-step/update-route-step
		// Handler: controllers.UpdateRouteStepHandler
		r.Put("/update-route-step", controllers.UpdateRouteStepHandler)

		// DELETE /route-step/delete-route-step/{route_name}/{stop_name}
		// Handler: controllers.DeleteRouteStepHandler
		r.Delete("/delete-route-step/{route_name}/{stop_name}", controllers.DeleteRouteStepHandler)
	})
}
