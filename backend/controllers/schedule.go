package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"
	"strconv"

	"github.com/go-chi/chi/v5"
	// "time"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	schedules, err := services.GetSchedule()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(schedules)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBusSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule structs.NewSchedule    
    err := json.NewDecoder(r.Body).Decode(&schedule)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = services.CreateBusSchedule(schedule)
    if err != nil {
        http.Error(w, "Failed to create schedule", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Schedule created successfully")
}

func UpdateBusSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule structs.UpdateSchedule    
    err := json.NewDecoder(r.Body).Decode(&schedule)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = services.UpdateBusSchedule(schedule)
    if err != nil {
        http.Error(w, "Failed to update schedule", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Schedule updated successfully")
}

func DeleteBusSchedule(w http.ResponseWriter, r *http.Request) {
    scheduleID := chi.URLParam(r, "id")
    if scheduleID == "" {
        http.Error(w, "Schedule ID is required", http.StatusBadRequest)
        return
    }

    scheduleIDInt, err := strconv.Atoi(scheduleID)
    if err != nil {
        http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
        return
    }

    err = services.DeleteBusSchedule(scheduleIDInt)
    if err != nil {
        http.Error(w, "Failed to delete schedule", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Schedule deleted successfully")
}