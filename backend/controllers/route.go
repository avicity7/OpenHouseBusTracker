package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"

	"github.com/go-chi/chi/v5"
)

// "github.com/jackc/pgx/v5"

// create
func CreateRouteHandler(w http.ResponseWriter, r *http.Request) {
	var route structs.Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := services.CreateRoute(route); err != nil {
		http.Error(w, "Failed to create route", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// for all routes
func GetAllRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := services.GetAllRoutes()
	if err != nil {
		fmt.Println(err)
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

// get requests to fetch a route by its name
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

// delete
func DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "route_name")
	if err := services.DeleteRoute(routeName); err != nil {
		http.Error(w, "Failed to delete route", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
