package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"server/services"
	"server/structs"
	"github.com/go-chi/chi/v5"
	"strconv"
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

	fmt.Fprintln(w, "Driver added successfully")
}

// Get Driver
func GetDriver(w http.ResponseWriter, r *http.Request) {
    drivers, err := services.GetDriver()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(drivers); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Update Driver
func UpdateDriver(w http.ResponseWriter, r *http.Request) {
    var driver structs.Driver
    err := json.NewDecoder(r.Body).Decode(&driver)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    idStr := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid driver ID", http.StatusBadRequest)
        return
    }
    drivers, err := services.GetDriver()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var currentDriver *structs.Driver
    for _, d := range drivers {
        if d.DriverId == int(id) {
            currentDriver = &d
            break
        }
    }
    if currentDriver == nil {
        http.Error(w, "Driver not found", http.StatusNotFound)
        return
    }
	if driver.DriverName == currentDriver.DriverName {
		w.WriteHeader(http.StatusNoContent)
		return
	}
    driver.DriverId = int(id)
    err = services.UpdateDriver(driver)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, "Driver updated successfully")
    w.WriteHeader(http.StatusOK)
}

// Delete Driver
func DeleteDriver(w http.ResponseWriter, r *http.Request) {
    driverID := chi.URLParam(r, "id") // Retrieve driverID from URL param

    id, err := strconv.Atoi(driverID)
    if err != nil {
        http.Error(w, "Invalid driver ID", http.StatusBadRequest)
        return
    }

    drivers, err := services.GetDriver()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var currentDriver *structs.Driver
    for _, d := range drivers {
        if d.DriverId == id {
            currentDriver = &d
            break
        }
    }

    if currentDriver == nil {
        http.Error(w, "Driver not found", http.StatusNotFound)
        return
    }

    err = services.DeleteDriver(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, "Driver deleted successfully")
}
