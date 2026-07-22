package dto

type CreateVehicleRequest struct {
	VIN string `json:"vin"`

	Year int `json:"year" binding:"required"`

	Make string `json:"make" binding:"required"`

	Model string `json:"model" binding:"required"`

	Trim string `json:"trim"`

	Engine string `json:"engine"`

	Drivetrain string `json:"drivetrain"`

	Transmission string `json:"transmission"`

	Mileage int `json:"mileage"`

	Color string `json:"color"`

	LicensePlate string `json:"license_plate"`

	State string `json:"state"`

	IsPrimary bool `json:"is_primary"`
}
