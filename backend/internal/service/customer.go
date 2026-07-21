package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type CustomerService struct {
	customers repository.CustomerRepository
}

func NewCustomerService(
	customers repository.CustomerRepository,
) *CustomerService {

	return &CustomerService{
		customers: customers,
	}

}

func (s *CustomerService) Get(
	ctx context.Context,
	userID uuid.UUID,
) (
	*customer.Customer,
	error,
) {

	return s.customers.FindByUserID(
		ctx,
		userID,
	)

}

func (s *CustomerService) Update(
	ctx context.Context,
	c *customer.Customer,
) error {

	return s.customers.Update(
		ctx,
		c,
	)

}
