package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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

func CreateEventHelpers(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventHelpers []structs.EventHelper
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Request body:", string(body))

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	eventHelpers := req.EventHelpers

	for _, eventHelper := range eventHelpers {
		email, err := services.GetEmail(eventHelper.Name)
		if err != nil {
			fmt.Println("Error retrieving email:", err)
			http.Error(w, "Failed to retrieve email", http.StatusInternalServerError)
			return
		}

		exists, err := services.EventHelperExists(email)
		if err != nil {
			http.Error(w, "Failed to check existing event helper", http.StatusInternalServerError)
			return
		}

		if exists {
			http.Error(w, "Student is already assigned to a carplate", http.StatusConflict)
			return
		}
	}

	fmt.Println("WHAT IS INSIDE THIS", eventHelpers)
	err = services.CreateEventHelpers(eventHelpers)
	if err != nil {
		fmt.Println("Error creating event helpers:", err)
		http.Error(w, "Failed to create event helpers", http.StatusInternalServerError)
		return
	}

	config.Cache.Delete("EventHelpers")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Event Helpers created successfully")
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
	fmt.Println("Values:")
	fmt.Println("NewCarplate:", eventHelper.NewBusId)
	fmt.Println("NewEmail:", eventHelper.NewName)
	fmt.Println("NewShift:", eventHelper.NewShift)
	fmt.Println("OldCarplate:", eventHelper.OldBusId)
	fmt.Println("OldEmail:", eventHelper.OldName)
	fmt.Println("OldShift:", eventHelper.OldShift)

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
