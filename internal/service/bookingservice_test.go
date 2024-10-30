// internal/service/bookingservice_test.go

package service

import (
	"context"
	"log"
	"os"
	"space-trouble/internal/model"
	"space-trouble/internal/repository"
	"space-trouble/pkg/db"
	"testing"
)

// TestBookingsService_CreateBooking tests the creation of a booking through the service layer.
func TestBookingService_CreateBooking(t *testing.T) {
	// Reading database connection string from environment variables
	dbConnString := os.Getenv("DB_CONN_STRING")
	if dbConnString == "" {
		log.Fatal("DB_CONN_STRING environment variable is not set")
	}

	dbConn, err := db.NewPostgresDB(dbConnString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	repo := repository.NewBookingRepository(dbConn)
	service := NewBookingService(repo)
	ctx := context.Background()

	// Valid booking scenario
	booking := &model.Booking{
		FirstName:     "John",
		LastName:      "Doe",
		Gender:        "Male",
		Birthday:      "1985-01-01",
		LaunchpadID:   "1",
		DestinationID: "Mars",
		LaunchDate:    "2049-12-31",
	}
	err = service.CreateBooking(ctx, booking)
	if err != nil {
		t.Errorf("Failed to create booking through service: %v", err)
	}
}

// TestBookingsService_GetAllBookings tests retrieval of all bookings through the service layer.
func TestBookingService_GetAllBookings(t *testing.T) {
	dbConn, err := db.NewPostgresDB("your_test_database_connection_string_here")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	repo := repository.NewBookingRepository(dbConn)
	service := NewBookingService(repo)
	ctx := context.Background()

	bookings, err := service.GetAllBookings(ctx)
	if err != nil {
		t.Errorf("Failed to get all bookings through service: %v", err)
	}
	if len(bookings) == 0 {
		t.Errorf("Expected one or more bookings, got %d", len(bookings))
	}
}

// TestBookingService_CreateBookingWithInvalidData tests creating a booking with invalid data.
func TestBookingService_CreateBookingWithInvalidData(t *testing.T) {
	dbConn, err := db.NewPostgresDB("your_test_database_connection_string_here")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	repo := repository.NewBookingRepository(dbConn)
	service := NewBookingService(repo)
	ctx := context.Background()

	// Invalid booking scenario: missing required fields
	booking := &model.Booking{
		FirstName:     "", // Missing first name
		LastName:      "Doe",
		Gender:        "Male",
		Birthday:      "1985-01-01",
		LaunchpadID:   "1",
		DestinationID: "Mars",
		LaunchDate:    "2049-12-31",
	}
	err = service.CreateBooking(ctx, booking)
	if err == nil {
		t.Errorf("Expected error when creating booking with invalid data, got none")
	}
}
