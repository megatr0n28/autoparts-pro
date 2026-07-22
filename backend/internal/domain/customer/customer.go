package customer

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (Customer) TableName() string {
	return "customer_profiles"
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {

	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return nil
}
