package controllers

import (
	"encoding/json"
	"net/http"
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
	carplate := chi.URLParam(r, "carplate")

	status, err := services.GetBusStatus(carplate)
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

func UpdateBusVisibility(w http.ResponseWriter, r *http.Request) {
	carplate := chi.URLParam(r, "carplate")
	hidden := chi.URLParam(r, "hidden")

	err := services.UpdateBusVisibility(hidden, carplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func DeleteBus(w http.ResponseWriter, r *http.Request) {
	carplate := chi.URLParam(r, "carplate")

	err := services.DeleteBus(carplate)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}
