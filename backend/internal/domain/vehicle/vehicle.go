package vehicle

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID uuid.UUID

	CustomerID uuid.UUID

	VIN string

	Year int

	Make string

	Model string

	Trim string

	Engine string

	Drivetrain string

	Transmission string

	Mileage int

	Color string

	LicensePlate string

	State string

	IsPrimary bool

	CreatedAt time.Time

	UpdatedAt time.Time
}

func (Vehicle) TableName() string {
	return "vehicles"
}

func (v *Vehicle) BeforeCreate(tx *gorm.DB) error {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}

	return nil
}
