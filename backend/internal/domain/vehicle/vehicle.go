package vehicle

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID uuid.UUID `json:"id"`

	CustomerID uuid.UUID `json:"customer_id"`

	VIN string `json:"vin"`

	Year int `json:"year"`

	Make string `json:"make"`

	Model string `json:"model"`

	Trim string `json:"trim"`

	Engine string `json:"engine"`

	Drivetrain string `json:"drivetrain"`

	Transmission string `json:"transmission"`

	Mileage int `json:"mileage"`

	Color string `json:"color"`

	LicensePlate string `json:"license_plate"`

	State string `json:"state"`

	IsPrimary bool `json:"is_primary"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
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
