package repository

import (
	"time"

	"github.com/conrad3rd/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUser() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}
