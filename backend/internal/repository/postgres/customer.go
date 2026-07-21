package postgres

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(
	db *gorm.DB,
) *CustomerRepository {

	return &CustomerRepository{
		db: db,
	}

}

func (r *CustomerRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (
	*customer.Customer,
	error,
) {

	var c customer.Customer

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"user_id = ?",
				userID,
			).
			First(&c).
			Error

	return &c, err

}

func (r *CustomerRepository) Create(
	ctx context.Context,
	c *customer.Customer,
) error {

	return r.db.
		WithContext(ctx).
		Create(c).
		Error

}

func (r *CustomerRepository) Update(
	ctx context.Context,
	c *customer.Customer,
) error {

	return r.db.
		WithContext(ctx).
		Save(c).
		Error

}
