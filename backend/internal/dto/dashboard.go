package dto

import vehicleDomain "github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"

type DashboardResponse struct {
	Customer     CustomerResponse        `json:"customer"`
	Vehicles     []vehicleDomain.Vehicle `json:"vehicles"`
	VehicleCount int                     `json:"vehicle_count"`
}
