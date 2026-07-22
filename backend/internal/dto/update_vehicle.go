package dto

type UpdateVehicleRequest struct {
	VIN string `json:"vin"`

	Year int `json:"year"`

	Make string `json:"make"`

	Model string `json:"model"`

	Trim string `json:"trim"`

	Engine string `json:"engine"`

	Transmission string `json:"transmission"`

	Drivetrain string `json:"drivetrain"`

	Mileage int `json:"mileage"`

	Color string `json:"color"`

	LicensePlate string `json:"license_plate"`

	State string `json:"state"`
}
