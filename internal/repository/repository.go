package repository

import "github.com/conrad3rd/bookings/internal/models"

type DatabaseRepo interface {
	AllUser() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.Restriction) error
}
