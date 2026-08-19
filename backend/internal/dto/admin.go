package dto

import (
	"time"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
)

type AdminUserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminCustomerResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
}

type AdminInvoiceResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type AdminOverviewResponse struct {
	Users                    []AdminUserResponse     `json:"users"`
	Customers                []AdminCustomerResponse `json:"customers"`
	Vehicles                 []vehicle.Vehicle       `json:"vehicles"`
	Invoices                 []AdminInvoiceResponse  `json:"invoices"`
	InvoiceManagementEnabled bool                    `json:"invoice_management_enabled"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" binding:"required"`
}
