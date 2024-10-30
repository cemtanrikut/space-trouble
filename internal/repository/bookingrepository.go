package repository

import (
	"context"
	"database/sql"

	"space-trouble/internal/model"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) CreateBooking(ctx context.Context, b *model.Booking) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		b.FirstName,
		b.LastName,
		b.Gender,
		b.Birthday,
		b.LaunchpadID,
		b.DestinationID,
		b.LaunchDate)
	return err
}

func (r *BookingRepository) GetAllBookings(ctx context.Context) ([]*model.Booking, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date FROM bookings")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookings []*model.Booking
	for rows.Next() {
		var b model.Booking
		if err := rows.Scan(
			&b.ID,
			&b.FirstName,
			&b.LastName,
			&b.Gender,
			&b.Birthday,
			&b.LaunchpadID,
			&b.DestinationID,
			&b.LaunchDate); err != nil {
			return nil, err
		}
		bookings = append(bookings, &b)
	}

	return bookings, nil
}
