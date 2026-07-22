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

func (r *VehicleRepository) Update(
	ctx context.Context,
	v *vehicle.Vehicle,
) error {

	return r.db.
		WithContext(ctx).
		Save(v).
		Error
}

func (r *VehicleRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where(
			"id = ? AND customer_id = ?",
			id,
			customerID,
		).
		Delete(&vehicle.Vehicle{}).
		Error
}

func (r *VehicleRepository) ClearPrimary(
	ctx context.Context,
	customerID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Model(&vehicle.Vehicle{}).
		Where(
			"customer_id = ?",
			customerID,
		).
		Update(
			"is_primary",
			false,
		).
		Error
}

func (r *VehicleRepository) SetPrimary(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Model(&vehicle.Vehicle{}).
		Where(
			"id = ? AND customer_id = ?",
			id,
			customerID,
		).
		Update(
			"is_primary",
			true,
		).
		Error
}

func (r *VehicleRepository) GetCustomerVehicle(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
) (*vehicle.Vehicle, error) {

	var v vehicle.Vehicle

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"id = ? AND customer_id = ?",
				id,
				customerID,
			).
			First(&v).
			Error

	return &v, err
}

func (r *VehicleRepository) FindByVIN(
	ctx context.Context,
	vin string,
) (*vehicle.Vehicle, error) {

	var v vehicle.Vehicle

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"vin = ?",
				vin,
			).
			First(&v).
			Error

	return &v, err
}
