package routes

import (
	"net/http"
	"server/controllers"

	"github.com/go-chi/chi/v5"
)

// Chat sets up the routes for the chat module.
func Chat(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		// CreateChatRoom creates a new chat room.
		r.Put("/create-chat-room", http.HandlerFunc(controllers.CreateChatRoom))

		// CreateMessage creates a new message in a chat room.
		r.Put("/create-message", http.HandlerFunc(controllers.CreateMessage))

		// GetChatRooms retrieves the chat rooms for a given email.
		r.Get("/get-chat-rooms/{email}", http.HandlerFunc(controllers.GetChatRooms))

		// GetMessages retrieves the messages for a given room ID.
		r.Get("/get-messages/{room_id}", http.HandlerFunc(controllers.GetMessages))

		// DeleteMessage deletes a message from a chat room.
		r.Delete("/delete-message", http.HandlerFunc(controllers.DeleteMessage))

		// DeleteRoom deletes a chat room.
		r.Delete("/delete-room/{room_id}", http.HandlerFunc(controllers.DeleteRoom))
	})
}
