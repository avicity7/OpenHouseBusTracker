package main

import (
	"net/http"
	"server/config"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	config.Connect(r)

	routes.TestRoute(r)
	routes.Auth(r)
	routes.Users(r)

	http.ListenAndServe("127.0.0.1:3000", r)
}
