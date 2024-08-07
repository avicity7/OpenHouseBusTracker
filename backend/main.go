// Module server is the entry point for the OpenHouseBusTracker backend application.
// Through package main it initializes the server, sets up routes, and starts the server.
package main

import (
	"net/http"
	"os"
	"server/config"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

// main is the entry point for the OpenHouseBusTracker backend application.
// It loads environment variables, initializes the router, sets up middleware,
// connects to the database, registers routes, and starts the server.
func main() {
	godotenv.Load()

	// Create a new router
	r := chi.NewRouter()

	// Configure CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://open-house-bus-tracker.vercel.app", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	// Add logger middleware
	r.Use(middleware.Logger)

	// Connect to the database
	config.Connect(r)

	// Register routes
	routes.TestRoute(r)
	routes.Auth(r)
	routes.Users(r)
	routes.Route(r)
	routes.RouteStep(r)
	routes.Schedules(r)
	routes.Driver(r)
	routes.Event(r)
	routes.Bus(r)
	routes.EventHelper(r)
	routes.Chat(r)

	// Get the environment variable
	env := os.Getenv("ENV")

	// Start the server
	if env == "PROD" {
		http.ListenAndServeTLS(":3000", "fullchain.pem", "privkey.pem", r)
	} else {
		http.ListenAndServe("127.0.0.1:3000", r)
	}
}
