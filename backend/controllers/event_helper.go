package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/config"
	"server/services"
	"server/structs"

	// "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5"
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

func BulkCreateEventHelpers(w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
        fmt.Println("Error getting form file:", err)
        http.Error(w, "Invalid file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    csvReader := csv.NewReader(file)
    var eventHelpers []structs.EventHelper

    if _, err := csvReader.Read(); err != nil {
        fmt.Println("Error skipping header row:", err)
        http.Error(w, "Failed to parse CSV", http.StatusInternalServerError)
        return
    }

    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error reading CSV record:", err)
            http.Error(w, "Failed to parse CSV", http.StatusInternalServerError)
            return
        }

        fmt.Println("CSV Record:", record)

        var shift bool
        switch record[2] {
        case "AM":
            shift = true
        case "PM":
            shift = false
        default:
            fmt.Println("Invalid shift value:", record[2])
            http.Error(w, "Invalid shift value", http.StatusBadRequest)
            return
        }

        busId, err := services.GetBusIdByCarplate(record[0])
        if err != nil {
            fmt.Println("Error getting bus ID:", err)
            http.Error(w, "Failed to get Bus ID", http.StatusInternalServerError)
            return
        }

        eventHelper := structs.EventHelper{
            BusId:    busId,
            Carplate: record[0],
            Email:    record[1],
            Shift:    shift,
        }
        eventHelpers = append(eventHelpers, eventHelper)
    }

    err = services.BulkCreateEventHelpers(eventHelpers)
    if err != nil {
        fmt.Println("Error creating event helpers:", err)
        http.Error(w, "Failed to create event helpers", http.StatusInternalServerError)
        return
    }
    config.Cache.Delete("EventHelpers")

    w.WriteHeader(http.StatusOK)
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

func GetAvailableSwaps(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	helpers, err := services.GetAvailableSwaps(email)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	response, err := json.Marshal(helpers)
	if err != nil {
		http.Error(w, "Failed to marshal dropdown data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func CreateSwapRequest(w http.ResponseWriter, r *http.Request) {
	var swap_request structs.SwapRequest
	err := json.NewDecoder(r.Body).Decode(&swap_request)
	if err != nil {
		http.Error(w, "Error on Decoding Follow Bus", 500)
		return
	}

	err = services.CreateSwapRequest(swap_request)
	if err != nil {
		http.Error(w, "Failed to marshal dropdown data", http.StatusInternalServerError)
	}

	w.WriteHeader(200)
}

func GetSwapRequests(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	swap_requests, err := services.GetSwapRequests(email)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	response, err := json.Marshal(swap_requests)
	if err != nil {
		http.Error(w, "Failed to marshal dropdown data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}

func AcceptSwapRequest(w http.ResponseWriter, r *http.Request) {
	var swap_request structs.SwapRequest
	err := json.NewDecoder(r.Body).Decode(&swap_request)
	if err != nil {
		http.Error(w, "Error on Decoding Follow Bus", 500)
		return
	}

	err = services.AcceptSwapRequest(swap_request)
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}

func DeleteSwapRequest(w http.ResponseWriter, r *http.Request) {
	var swap_request structs.SwapRequest
	err := json.NewDecoder(r.Body).Decode(&swap_request)
	if err != nil {
		http.Error(w, "Error on Decoding Follow Bus", 500)
		return
	}

	err = services.DeleteSwapRequest(swap_request)
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}
