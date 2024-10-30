package service

import (
	"context"
	"space-trouble/internal/model"
	"space-trouble/internal/repository"
)

type BookingService struct {
	repo *repository.BookingRepository
}

func NewBookingService(repo *repository.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) CreateBooking(ctx context.Context, b *model.Booking) error {
	return s.repo.CreateBooking(ctx, b)
}

func (s *BookingService) GetAllBookings(ctx context.Context) ([]*model.Booking, error) {
	return s.repo.GetAllBookings(ctx)
}
