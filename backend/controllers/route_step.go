package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"
	"server/structs"

	"github.com/go-chi/chi/v5"
)

func GetAllRouteSteps(w http.ResponseWriter, r *http.Request) {
	routeSteps, err := services.GetAllRouteSteps()
	if err != nil {
		http.Error(w, "Failed to fetch route steps", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(routeSteps)
}

func GetRouteSteps(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "routeName")
	routeSteps, err := services.GetRouteSteps(routeName)
	if err != nil {
		http.Error(w, "Failed to fetch route steps", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(routeSteps)
}

func GetAllStopNames(w http.ResponseWriter, r *http.Request) {
	routeSteps, err := services.GetAllStops()
	if err != nil {
		http.Error(w, "Failed to fetch route steps", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(routeSteps)
}

// get
func GetRouteStepHandler(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "route_name")
	stopName := chi.URLParam(r, "stop_name")
	routeStep, err := services.GetRouteStep(routeName, stopName)
	if err != nil {
		http.Error(w, "Failed to fetch route step", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(routeStep)
}

// create
func CreateRouteStepHandler(w http.ResponseWriter, r *http.Request) {
	var routeStep structs.RouteStep
	if err := json.NewDecoder(r.Body).Decode(&routeStep); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := services.CreateRouteStep(routeStep); err != nil {
		http.Error(w, "Failed to create route step", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// update
func UpdateRouteStepHandler(w http.ResponseWriter, r *http.Request) {
	var routeStep structs.RouteStep
	if err := json.NewDecoder(r.Body).Decode(&routeStep); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := services.UpdateRouteStep(routeStep); err != nil {
		http.Error(w, "Failed to update route step", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// delete
func DeleteRouteStepHandler(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "route_name")
	stopName := chi.URLParam(r, "stop_name")
	if err := services.DeleteRouteStep(routeName, stopName); err != nil {
		http.Error(w, "Failed to delete route step", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
