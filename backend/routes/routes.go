package routes

import (
	"github.com/BhautikVekariya21/go/backend/controllers" // Full import path
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/bookings", controllers.BookingHandler).Methods("POST")
	router.HandleFunc("/api/bookings", controllers.GetBookingsHandler).Methods("GET")
	router.HandleFunc("/api/comments", controllers.CommentHandler).Methods("POST")
	router.HandleFunc("/api/comments", controllers.GetCommentsHandler).Methods("GET")
	router.HandleFunc("/api/contacts", controllers.CreateContact).Methods("POST") // Added route for contact creation
}
