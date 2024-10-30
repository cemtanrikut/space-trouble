package main

import (
	"log"
	"net/http"
	"os"
	"space-trouble/internal/app"
	"space-trouble/internal/repository"
	"space-trouble/internal/service"
	"space-trouble/pkg/db"
)

func main() {
	// Reading database connection string from environment variables
	dbConnString := os.Getenv("DB_CONN_STRING")
	if dbConnString == "" {
		log.Fatal("DB_CONN_STRING environment variable is not set")
	}

	// Database connection
	dbConn, err := db.NewPostgresDB(dbConnString)
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	defer dbConn.Close()

	// Setup repositories
	bookingRepo := repository.NewBookingRepository(dbConn)

	// Setup services
	bookingService := service.NewBookingService(bookingRepo)

	// Setup app
	myApp := app.App{Services: bookingService}

	// Setup router and start server
	router := myApp.SetupRouter()
	log.Println("Running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Error when starting server: ", err)
	}

}
