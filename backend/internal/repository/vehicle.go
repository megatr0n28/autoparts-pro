package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
)

type VehicleRepository interface {
	Create(
		ctx context.Context,
		v *vehicle.Vehicle,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*vehicle.Vehicle, error)

	GetByCustomer(
		ctx context.Context,
		customerID uuid.UUID,
	) ([]vehicle.Vehicle, error)

	Update(
		ctx context.Context,
		v *vehicle.Vehicle,
	) error

	Delete(
		ctx context.Context,
		id uuid.UUID,
		customerID uuid.UUID,
	) error

	SetPrimary(
		ctx context.Context,
		id uuid.UUID,
		customerID uuid.UUID,
	) error

	ClearPrimary(
		ctx context.Context,
		customerID uuid.UUID,
	) error
}
