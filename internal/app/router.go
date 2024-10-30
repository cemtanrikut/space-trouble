package app

import (
	"github.com/gorilla/mux"
)

func (app *App) SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Endpoint to create a booking
	router.HandleFunc("/bookings", app.bookTicketHandler).Methods("POST")

	// Endpoint to get all bookings
	router.HandleFunc("/bookings", app.getAllBookingsHandler).Methods("GET")

	spacexHandler := NewSpaceXHandler(app.SpaceXService)

	router.HandleFunc("/api/spacex/launches/upcoming", spacexHandler.GetUpcomingLaunches).Methods("GET")
	router.HandleFunc("/api/spacex/launchpads", spacexHandler.GetLaunchpads).Methods("GET")

	return router
}
