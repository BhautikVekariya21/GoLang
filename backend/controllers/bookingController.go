package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BhautikVekariya21/go/backend/models"   // Full import path
	"github.com/BhautikVekariya21/go/backend/services" // Full import path
)

// BookingHandler handles the creation of new bookings
func BookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateBooking(booking); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// GetBookingsHandler retrieves all bookings
func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := services.GetAllBookings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookings)
}
