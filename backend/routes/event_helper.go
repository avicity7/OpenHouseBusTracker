package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// EventHelper sets up the routes for event helper related endpoints.
func EventHelper(r chi.Router) {
	// /event-helpers
	r.Route("/event-helpers", func(r chi.Router) {
		// GET /event-helpers/get-accepted-swaps
		r.Get("/get-accepted-swaps", http.HandlerFunc(controllers.GetAcceptedSwapRequests))

		// GET /event-helpers/get-available-swaps/{email}
		r.Get("/get-available-swaps/{email}", http.HandlerFunc(controllers.GetAvailableSwaps))

		// GET /event-helpers/get-swap-requests/{email}
		r.Get("/get-swap-requests/{email}", http.HandlerFunc(controllers.GetSwapRequests))

		// GET /event-helpers/get-helpers
		r.Get("/get-helpers", http.HandlerFunc(controllers.GetEventHelpers))

		// GET /event-helpers/get-event-dropdown
		r.Get("/get-event-dropdown", http.HandlerFunc(controllers.GetEventHelperDropdownData))

		// POST /event-helpers/create-helpers
		r.Post("/create-helpers", http.HandlerFunc(controllers.CreateEventHelpers))

		// POST /event-helpers/bulk-create-shifts
		r.Post("/bulk-create-shifts", http.HandlerFunc(controllers.BulkCreateEventHelpers))

		// POST /event-helpers/create-swap-request
		r.Post("/create-swap-request", http.HandlerFunc(controllers.CreateSwapRequest))

		// PUT /event-helpers/update-helper
		r.Put("/update-helper", http.HandlerFunc(controllers.UpdateEventHelper))

		// PUT /event-helpers/accept-swap-request
		r.Put("/accept-swap-request", http.HandlerFunc(controllers.AcceptSwapRequest))

		// DELETE /event-helpers/delete-helper
		r.Delete("/delete-helper", http.HandlerFunc(controllers.DeleteEventHelper))

		// DELETE /event-helpers/delete-swap-request
		r.Delete("/delete-swap-request", http.HandlerFunc(controllers.DeleteSwapRequest))
	})
}
