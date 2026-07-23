package dto

type PartSearchRequest struct {
	VehicleID string `form:"vehicle_id" binding:"required"`

	Query string `form:"query" binding:"required"`
}

type PartSearchResponse struct {
	Retailer string `json:"retailer"`

	Brand string `json:"brand"`

	PartNumber string `json:"part_number"`

	Name string `json:"name"`

	Description string `json:"description"`

	Price float64 `json:"price"`

	Currency string `json:"currency"`

	InStock bool `json:"in_stock"`

	ProductURL string `json:"product_url"`

	ImageURL string `json:"image_url"`
}
