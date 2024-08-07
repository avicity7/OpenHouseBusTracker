package routes

import (
	// "github.com/jackc/pgx/v5"
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Route sets up the routes for handling route-related operations.
//
// Available routes:
//   - GET /route: Retrieves all routes.
//   - GET /route/{name}: Retrieves a route by its name.
//   - POST /route/create-route: Creates a new route.
//   - DELETE /route/delete-route/{route_name}: Deletes a route by its name.
func Route(r *chi.Mux) {
	r.Route("/route", func(r chi.Router) {
		r.Get("/", http.HandlerFunc(controllers.GetAllRoutesHandler))
		r.Get("/{name}", http.HandlerFunc(controllers.GetRouteByNameHandler))
		r.Post("/create-route", controllers.CreateRouteHandler)
		r.Delete("/delete-route/{route_name}", controllers.DeleteRouteHandler)
	})
}
