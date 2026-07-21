package dto

type CustomerResponse struct {
	ID string `json:"id"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Phone string `json:"phone"`

	AddressLine1 string `json:"address_line1"`

	AddressLine2 string `json:"address_line2"`

	City string `json:"city"`

	State string `json:"state"`

	PostalCode string `json:"postal_code"`

	Country string `json:"country"`
}

type UpdateCustomerRequest struct {
	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Phone string `json:"phone"`

	AddressLine1 string `json:"address_line1"`

	AddressLine2 string `json:"address_line2"`

	City string `json:"city"`

	State string `json:"state"`

	PostalCode string `json:"postal_code"`
}
