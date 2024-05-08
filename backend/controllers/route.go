package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"

	"github.com/go-chi/chi/v5"
)

// "github.com/jackc/pgx/v5"

// for all routes
func GetAllRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := services.GetAllRoutes()
	if err != nil {
		http.Error(w, "Failed to fetch routes", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(routes)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GET requests to fetch a route by its name
func GetRouteByNameHandler(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "name")

	if routeName == "" {
		http.Error(w, "Route name is required", http.StatusBadRequest)
		return
	}

	route, err := services.GetRouteByName(routeName)
	if err != nil {
		http.Error(w, "Failed to fetch route", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(route)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
