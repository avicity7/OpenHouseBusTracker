package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

func Chat(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		r.Put("/create-chat-room", http.HandlerFunc(controllers.CreateChatRoom))
		r.Put("/create-message", http.HandlerFunc(controllers.CreateMessage))
		r.Get("/get-chat-rooms/{email}", http.HandlerFunc(controllers.GetChatRooms))
		r.Get("/get-messages/{room_id}", http.HandlerFunc(controllers.GetMessages))
		r.Delete("/delete-message", http.HandlerFunc(controllers.DeleteMessage))
		r.Delete("/delete-room/{room_id}", http.HandlerFunc(controllers.DeleteRoom))
	})
}
