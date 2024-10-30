package app

import (
	"encoding/json"
	"net/http"
	"space-trouble/internal/model"
	"space-trouble/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Services      *service.BookingService
	SpaceXService *service.SpaceXService
}

func (app *App) bookTicketHandler(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.Services.CreateBooking(r.Context(), &booking); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *App) getAllBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := app.Services.GetAllBookings(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookings)
}

// deleteBookingHandler deletes a booking by ID
func (app *App) deleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	if err := app.Services.DeleteBooking(r.Context(), int64(id)); err != nil {
		http.Error(w, "Could not delete booking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
