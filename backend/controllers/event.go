package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateFollowBus(w http.ResponseWriter, r *http.Request) {
	var followBus structs.FollowBus

	err := json.NewDecoder(r.Body).Decode(&followBus)
	if err != nil {
		http.Error(w, "Error on Decoding Follow Bus", 500)
		return
	}

	err = services.CreateFollowBus(followBus.Carplate, followBus.Email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on Creating Follow Bus", 500)
		return
	}

	w.WriteHeader(200)
}

func GetFollowBus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	response, err := services.GetFollowBus(email)
	if err != nil {
		http.Error(w, "Follow Bus not found", 500)
		return
	}

	parsed, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marhsalling repsonse", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func GetAllFollowBus(w http.ResponseWriter, r *http.Request) {
	response, err := services.GetAllFollowBus()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on Service", 500)
	}

	parsed, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error on marshalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func DeleteFollowBus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	err := services.DeleteFollowBus(email)
	if err != nil {
		http.Error(w, "Error on Deleting Follow Bus", 500)
		return
	}

	w.WriteHeader(200)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	scheduleId := chi.URLParam(r, "scheduleId")
	scheduleIdInt, _ := strconv.Atoi(scheduleId)

	response, err := services.GetEvents(scheduleIdInt)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on service", 500)
		return
	}

	parsed, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on marshalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input structs.EventInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on Decoding", 500)
		return
	}

	err = services.CreateEvent(input.BusId, input.RouteName, input.EventId, input.StopName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on Service", 500)
		return
	}

	var status bool

	if input.EventId == 4 || input.EventId == 5 {
		status = false
	} else if input.EventId == 1 || input.EventId == 2 {
		status = true
	}

	err = services.UpdateBusStatus(status, input.BusId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on update bus", 500)
		return
	}

	w.WriteHeader(200)
}

func DeleteLastEvent(w http.ResponseWriter, r *http.Request) {
	bus_id := chi.URLParam(r, "bus_id")
	err := services.DeleteLastEvent(bus_id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func GetAllStops(w http.ResponseWriter, r *http.Request) {
	output, err := services.GetAllStops()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting routes", 500)
		return
	}

	parsed, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error marhsalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func GetCurrentBuses(w http.ResponseWriter, r *http.Request) {
	currentBuses, err := services.GetCurrentBuses()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	formatted, err := json.Marshal(currentBuses)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(formatted)
}
