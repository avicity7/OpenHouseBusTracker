package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"
	"server/structs"
)

// Add Driver
func AddDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AddDriver(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get Driver
func GetDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverID := vars["id"]

	driver, err := services.GetDriver(driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(driver)
}

// Update Driver
func UpdateDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateDriver(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Driver
func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverID := vars["id"]

	err := services.DeleteDriver(driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
