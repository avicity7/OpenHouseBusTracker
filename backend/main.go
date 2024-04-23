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

	env := os.Getenv("ENV")
	if env == "PROD" {
		http.ListenAndServeTLS("127.0.0.1:3000", "./cert.pem", "./key.pem", r)
	} else {
		http.ListenAndServe("127.0.0.1:3000", r)
	}
}
