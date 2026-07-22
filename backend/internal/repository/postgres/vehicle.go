package postgres

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
)

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(
	db *gorm.DB,
) *VehicleRepository {

	return &VehicleRepository{
		db: db,
	}
}

func (r *VehicleRepository) Create(
	ctx context.Context,
	v *vehicle.Vehicle,
) error {

	return r.db.
		WithContext(ctx).
		Create(v).
		Error
}

func (r *VehicleRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*vehicle.Vehicle, error) {

	var v vehicle.Vehicle

	err :=
		r.db.
			WithContext(ctx).
			First(
				&v,
				"id = ?",
				id,
			).Error

	return &v, err
}

func (r *VehicleRepository) GetByCustomer(
	ctx context.Context,
	customerID uuid.UUID,
) ([]vehicle.Vehicle, error) {

	var vehicles []vehicle.Vehicle

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"customer_id = ?",
				customerID,
			).
			Order("created_at DESC").
			Find(&vehicles).
			Error

	return vehicles, err
}
