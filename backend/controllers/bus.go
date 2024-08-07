package controllers

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/services"

	"github.com/go-chi/chi/v5"
)

func CreateBus(w http.ResponseWriter, r *http.Request) {
	carplate := chi.URLParam(r, "carplate")

	err := services.CreateBus(carplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func GetBuses(w http.ResponseWriter, r *http.Request) {
	err := services.RefreshBuses()
	if err != nil {
		http.Error(w, "Error getting Buses", 500)
		return
	}
	response, err := services.GetBuses()
	if err != nil {
		http.Error(w, "Error getting Buses", 500)
		return
	}

	parsed, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error on marshalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func GetBusStatus(w http.ResponseWriter, r *http.Request) {
	err := services.RefreshBuses()
	if err != nil {
		http.Error(w, "Error getting Buses", 500)
		return
	}
	bus_id := chi.URLParam(r, "bus_id")

	status, err := services.GetBusStatus(bus_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parsed, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func UpdateBus(w http.ResponseWriter, r *http.Request) {
	bus_id := chi.URLParam(r, "bus_id")
	carplate := chi.URLParam(r, "carplate")

	err := services.UpdateBus(bus_id, carplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	config.Cache.Delete("Schedules")

	w.WriteHeader(200)
}

func UpdateBusVisibility(w http.ResponseWriter, r *http.Request) {
	bus_id := chi.URLParam(r, "bus_id")
	hidden := chi.URLParam(r, "hidden")

	err := services.UpdateBusVisibility(hidden, bus_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func DeleteBus(w http.ResponseWriter, r *http.Request) {
	bus_id := chi.URLParam(r, "bus_id")

	err := services.DeleteBus(bus_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	config.Cache.Delete("Schedules")

	w.WriteHeader(200)
}
