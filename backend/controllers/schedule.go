package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"
	"server/config"	
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
    value, found := config.Cache.Get("Schedules")
    if found {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(value.([]byte))
        return
    }

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

    config.Cache.Set("Schedules", response, cache.DefaultExpiration)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func CreateBusSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule structs.NewSchedule

    err := json.NewDecoder(r.Body).Decode(&schedule)

    if err != nil {
        fmt.Println(err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = services.CreateBusSchedule(schedule)
    if err != nil {
        http.Error(w, "Failed to create schedule", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("Schedules")
    config.Cache.Delete("CurrentUserSchedules")
    config.Cache.Delete("FutureUserSchedules")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Schedule created successfully")
}

func UpdateBusSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule structs.UpdateSchedule    
    err := json.NewDecoder(r.Body).Decode(&schedule)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        fmt.Println("Error decoding request body:", err)
        return
    }

    fmt.Printf("Received update request for schedule: %+v\n", schedule)

    err = services.UpdateBusSchedule(schedule)
    if err != nil {
        http.Error(w, "Failed to update schedule", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("Schedules")
    config.Cache.Delete("CurrentUserSchedules")
    config.Cache.Delete("FutureUserSchedules")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Schedule updated successfully")
}

func DeleteBusSchedule(w http.ResponseWriter, r *http.Request) {
    // scheduleID := chi.URLParam(r, "id")
    schedule := make([]int, 0)
    err := json.NewDecoder(r.Body).Decode(&schedule)
    if len(schedule) == 0 {
        http.Error(w, "Schedule ID is required", http.StatusBadRequest)
        return
    }
    err = services.DeleteBusSchedule(schedule)
    if err != nil {
        http.Error(w, "Failed to delete schedule", http.StatusInternalServerError)
        return
    }

    config.Cache.Delete("Schedules")
    config.Cache.Delete("CurrentUserSchedules")
    config.Cache.Delete("FutureUserSchedules")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Schedule deleted successfully")
}


func GetDropdownData(w http.ResponseWriter, r *http.Request) {
	dropdownData, err := services.GetDropdownData()

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

func GetScheduleByID(w http.ResponseWriter, r *http.Request) {
	scheduleIDStr := chi.URLParam(r, "id")
	scheduleID, err := strconv.Atoi(scheduleIDStr)
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

    fmt.Println("Received schedule ID:", scheduleID) 

	schedule, err := services.GetScheduleByID(scheduleID)
	if err != nil {
		http.Error(w, "Failed to fetch schedule details", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(schedule)
	if err != nil {
		http.Error(w, "Failed to marshal schedule details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetScheduleByUser(w http.ResponseWriter, r *http.Request) {
    email := chi.URLParam(r, "email")
    if email == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    value, found := config.Cache.Get("CurrentUserSchedules:" + email)
    if found {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(value.([]byte))
        return
    }

    schedules, err := services.GetScheduleByUser(email)
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

    config.Cache.Set("CurrentUserSchedules:"+email, response, cache.DefaultExpiration)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

// might have caching issues right now
func GetFutureScheduleByUser(w http.ResponseWriter, r *http.Request) {
    email := chi.URLParam(r, "email")
    if email == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    value, found := config.Cache.Get("FutureUserSchedules:" + email)
    if found {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(value.([]byte))
        return
    }

    schedules, err := services.GetFutureScheduleByUser(email)
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

    config.Cache.Set("FutureUserSchedules:"+email, response, cache.DefaultExpiration)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}
