package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"
	"server/structs"

	"github.com/go-chi/chi/v5"
)

func CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	var chat_room structs.CreateChatRoom
	err := json.NewDecoder(r.Body).Decode(&chat_room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room_id, err := services.CreateChatRoom(chat_room.User1, chat_room.User2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parsed, err := json.Marshal(room_id)
	if err != nil {
		http.Error(w, "Error on marshalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func GetChatRooms(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	chat_rooms, err := services.GetChatRooms(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parsed, err := json.Marshal(chat_rooms)
	if err != nil {
		http.Error(w, "Error on marshalling", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message structs.CreateMessage

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateMessage(message.From, message.RoomId, message.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	room_id := chi.URLParam(r, "room_id")

	messages, err := services.GetMessages(room_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parsed, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(parsed)
}
