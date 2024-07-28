package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/services"
	"server/structs"

	"github.com/patrickmn/go-cache"
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
	config.Cache.Delete("drivers")
	fmt.Fprintln(w, "Driver added successfully")
}

// Get Driver
func GetDrivers(w http.ResponseWriter, r *http.Request) {
	if data, found := config.Cache.Get("drivers"); found {
		if drivers, ok := data.([]structs.Driver); ok {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(drivers); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	drivers, err := services.GetDrivers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	config.Cache.Set("drivers", drivers, cache.DefaultExpiration)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(drivers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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

	config.Cache.Delete("drivers")
	config.Cache.Delete("Schedules")

	w.WriteHeader(200)
}

func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	var driver structs.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteDriver(driver.DriverId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.Cache.Delete("drivers")
	w.WriteHeader(200)
}

func GetScheduleTimeDiff(w http.ResponseWriter, r *http.Request) {
	timeDiffs, err := services.GetScheduleTimeDiff()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeDiffs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
