package service

import (
	"context"
	"errors"
	"space-trouble/internal/model"
	"space-trouble/internal/repository"
	"space-trouble/internal/spacex"
	"time"
)

type BookingService struct {
	repo         *repository.BookingRepository
	spaceXClient *spacex.Client
}

func NewBookingService(repo *repository.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) GetAllBookings(ctx context.Context) ([]*model.Booking, error) {
	return s.repo.GetAllBookings(ctx)
}

// CheckLaunchAvailability checks if there is any SpaceX launch on the given date and launchpad
func (s *BookingService) CheckLaunchAvailability(ctx context.Context, launchpadID string, launchDate time.Time) (bool, error) {
	launches, err := s.spaceXClient.GetUpcomingLaunches()
	if err != nil {
		return false, err
	}

	for _, launch := range launches {
		launchTime, err := time.Parse(time.RFC3339, launch.LaunchDate)
		if err != nil {
			continue
		}

		// Same date and launchpad check
		if launch.LaunchpadID == launchpadID && launchTime.Equal(launchDate) {
			return false, nil
		}
	}
	return true, nil
}

// CreateBooking creates a booking if there is no SpaceX launch conflict
func (s *BookingService) CreateBooking(ctx context.Context, b *model.Booking) error {
	// Parse launch date
	launchDate, err := time.Parse("2006-01-02", b.LaunchDate)
	if err != nil {
		return err
	}

	// Check availability
	available, err := s.CheckLaunchAvailability(ctx, b.LaunchpadID, launchDate)
	if err != nil {
		return err
	}
	if !available {
		return errors.New("launch conflict: a SpaceX launch is scheduled for this date and launchpad")
	}

	return s.repo.CreateBooking(ctx, b)
}
