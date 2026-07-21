package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
)

type CustomerRepository interface {
	FindByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (
		*customer.Customer,
		error,
	)

	Create(
		ctx context.Context,
		customer *customer.Customer,
	) error

	Update(
		ctx context.Context,
		customer *customer.Customer,
	) error
}
