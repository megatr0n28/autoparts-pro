package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID uuid.UUID

	UserID uuid.UUID

	FirstName string

	LastName string

	Phone string

	AddressLine1 string

	AddressLine2 string

	City string

	State string

	PostalCode string

	Country string

	CreatedAt time.Time

	UpdatedAt time.Time
}
