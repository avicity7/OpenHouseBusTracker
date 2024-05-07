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

func DeleteFollowBus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	err := services.DeleteFollowBus(email)
	if err != nil {
		http.Error(w, "Error on Deleting Follow Bus", 500)
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

func GetRouteSteps(w http.ResponseWriter, r *http.Request) {
	routeName := chi.URLParam(r, "routeName")

	response, err := services.GetRouteSteps(routeName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on service", 500)
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

	err = services.CreateEvent(input.Carplate, input.RouteName, input.EventId, input.StopName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error on Service", 500)
		return
	}

	w.WriteHeader(200)
}
