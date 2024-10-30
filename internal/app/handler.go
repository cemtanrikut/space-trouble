package app

import (
	"encoding/json"
	"net/http"
	"space-trouble/internal/model"
	"space-trouble/internal/service"
)

type App struct {
	Services *service.BookingService
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
