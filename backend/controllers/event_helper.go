package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/services"
	"server/structs"

	// "github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
)

func GetEventHelpers(w http.ResponseWriter, r *http.Request) {
	value, found := config.Cache.Get("EventHelpers")
	if found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(value.([]byte))
		return
	}
    eventHelpers, err := services.GetEventHelpers()
    if err != nil {
		fmt.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    response, err := json.Marshal(eventHelpers)
    if err != nil {
		fmt.Println(err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	config.Cache.Set("EventHelpers", response, cache.DefaultExpiration)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func CreateEventHelper(w http.ResponseWriter, r *http.Request) {
    var eventHelper structs.EventHelper

    err := json.NewDecoder(r.Body).Decode(&eventHelper)

    if err != nil {
        fmt.Println(err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = services.CreateEventHelper(eventHelper)
    if err != nil {
        http.Error(w, "Failed to create event helper", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("EventHelpers")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Event Helper created successfully")
}


func UpdateEventHelper(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Received request body:", r.Body)
    var eventHelper structs.EventHelperUpdate    
    err := json.NewDecoder(r.Body).Decode(&eventHelper)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        fmt.Println("Request body:", r.Body)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        fmt.Println("Error decoding request body:", err)
        return
    }

    fmt.Printf("Received update request for event helper: %+v\n", eventHelper)

    fmt.Println("Executing SQL query:")
    fmt.Println("UPDATE event_helper SET carplate = $1, email = $2, start_time = $3, end_time = $4 WHERE carplate = $5 AND email = $6 AND start_time = $7 AND end_time = $8")
    fmt.Println("Values:")
    fmt.Println("NewCarplate:", eventHelper.NewCarplate)
    fmt.Println("NewEmail:", eventHelper.NewEmail)
    fmt.Println("NewStartTime:", eventHelper.NewStartTime)
    fmt.Println("NewEndTime:", eventHelper.NewEndTime)
    fmt.Println("OldCarplate:", eventHelper.OldCarplate)
    fmt.Println("OldEmail:", eventHelper.OldEmail)
    fmt.Println("OldStartTime:", eventHelper.OldStartTime)
    fmt.Println("OldEndTime:", eventHelper.OldEndTime)

    err = services.UpdateEventHelper(eventHelper)
    if err != nil {
        http.Error(w, "Failed to update Event Helper", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("EventHelpers")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Event Helpers updated successfully")
}


func DeleteEventHelper(w http.ResponseWriter, r *http.Request) {
    var eventHelper structs.EventHelper
    err := json.NewDecoder(r.Body).Decode(&eventHelper)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        fmt.Println("Error decoding request body:", err)
        return
    }

    err = services.DeleteEventHelper(eventHelper)
    if err != nil {
        http.Error(w, "Failed to delete Event Helpers", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("EventHelpers")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Event Helpers deleted successfully")
}

func GetEventHelperDropdownData(w http.ResponseWriter, r *http.Request) {
	dropdownData, err := services.GetEventHelperDropdownData()

	if err != nil {
		http.Error(w, "Failed to fetch dropdown data", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(dropdownData)
	if err != nil {
		http.Error(w, "Failed to marshal dropdown data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}