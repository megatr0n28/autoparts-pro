package testutil

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type MockCustomerRepository struct {
	customers map[uuid.UUID]*customer.Customer
}

func NewCustomerRepository() repository.CustomerRepository {

	return &MockCustomerRepository{
		customers: make(
			map[uuid.UUID]*customer.Customer,
		),
	}
}

func (m *MockCustomerRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*customer.Customer, error) {

	// Return existing customer if test created one
	for _, c := range m.customers {

		if c.UserID == userID {
			return c, nil
		}
	}

	// Default customer for authenticated smoke tests
	customerProfile :=
		&customer.Customer{

			ID: uuid.New(),

			UserID: userID,

			FirstName: "Test",

			LastName: "User",
		}

	m.customers[customerProfile.ID] =
		customerProfile

	return customerProfile, nil
}

func (m *MockCustomerRepository) Create(
	ctx context.Context,
	c *customer.Customer,
) error {

	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	m.customers[c.ID] = c

	return nil
}

func (m *MockCustomerRepository) Update(
	ctx context.Context,
	c *customer.Customer,
) error {

	m.customers[c.ID] = c

	return nil
}
