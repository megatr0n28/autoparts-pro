package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/authz"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
	vehicleService "github.com/megatr0n28/autoparts-pro/backend/internal/service/vehicle"
)

var ErrDashboardAccessDenied = errors.New("dashboard access denied")

type DashboardData struct {
	Customer *customer.Customer
	Vehicles []vehicle.Vehicle
}

type DashboardService struct {
	customers  *CustomerService
	vehicles   *vehicleService.Service
	authorizer *authz.Authorizer
}

func NewDashboardService(
	customers *CustomerService,
	vehicles *vehicleService.Service,
	authorizer *authz.Authorizer,
) *DashboardService {
	return &DashboardService{
		customers:  customers,
		vehicles:   vehicles,
		authorizer: authorizer,
	}
}

func (s *DashboardService) Get(
	ctx context.Context,
	userID string,
	customerID uuid.UUID,
) (*DashboardData, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid authenticated user ID: %w", err)
	}

	allowed, err := s.authorizer.Check(
		ctx,
		authz.UserObject(userUUID),
		authz.RelationOwner,
		authz.CustomerObject(customerID),
	)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, ErrDashboardAccessDenied
	}
	customerProfile, err := s.customers.Get(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	vehicles, err := s.vehicles.List(ctx, userID, customerID)
	if err != nil {
		return nil, err
	}

	return &DashboardData{
		Customer: customerProfile,
		Vehicles: vehicles,
	}, nil
}
