package repository

import (
	"context"
	"log"
	"os"
	"space-trouble/internal/model"
	"space-trouble/pkg/db"
	"testing"
)

var testRepo *BookingRepository

func TestMain(m *testing.M) {
	// Reading database connection string from environment variables
	dbConnString := os.Getenv("DB_CONN_STRING")
	if dbConnString == "" {
		log.Fatal("DB_CONN_STRING environment variable is not set")
	}

	// Database connection setup for testing
	dbConn, err := db.NewPostgresDB(dbConnString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	testRepo = NewBookingRepository(dbConn)

	// Run tests
	os.Exit(m.Run())
}

func TestCreateBooking(t *testing.T) {
	ctx := context.Background()
	booking := &model.Booking{
		FirstName:     "Test",
		LastName:      "User",
		Gender:        "Other",
		Birthday:      "2000-01-01",
		LaunchpadID:   "1",
		DestinationID: "Mars",
		LaunchDate:    "2049-01-01",
	}
	err := testRepo.CreateBooking(ctx, booking)
	if err != nil {
		t.Errorf("Failed to create booking: %v", err)
	}

	// Optionally check if the booking is actually added by retrieving it
}

func TestGetAllBookings(t *testing.T) {
	ctx := context.Background()
	bookings, err := testRepo.GetAllBookings(ctx)
	if err != nil {
		t.Errorf("Failed to get all bookings: %v", err)
	}
	if len(bookings) == 0 {
		t.Errorf("Expected one or more bookings, got %d", len(bookings))
	}
}

func TestCreateBookingWithInvalidData(t *testing.T) {
	ctx := context.Background()
	booking := &model.Booking{
		// Intentionally leaving some fields empty to simulate bad input
		FirstName:   "",
		LastName:    "User",
		Gender:      "Other",
		Birthday:    "2000-01-01",
		LaunchpadID: "1", // Assume this is a valid ID
	}
	err := testRepo.CreateBooking(ctx, booking)
	if err == nil {
		t.Errorf("Expected error when creating booking with invalid data, got none")
	}
}
