package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type AdminOverview struct {
	Users     []user.User
	Customers []customer.Customer
	Vehicles  []vehicle.Vehicle
}

type AdminService struct {
	users     repository.UserRepository
	customers repository.CustomerRepository
	vehicles  repository.VehicleRepository
}

func NewAdminService(
	users repository.UserRepository,
	customers repository.CustomerRepository,
	vehicles repository.VehicleRepository,
) *AdminService {
	return &AdminService{users: users, customers: customers, vehicles: vehicles}
}

func (s *AdminService) Overview(ctx context.Context) (*AdminOverview, error) {
	users, err := s.users.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	customers, err := s.customers.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	vehicles, err := s.vehicles.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &AdminOverview{Users: users, Customers: customers, Vehicles: vehicles}, nil
}

func (s *AdminService) UpdateUserRole(
	ctx context.Context,
	id uuid.UUID,
	role string,
) error {
	if role != "viewer" && role != "editor" && role != "admin" {
		return fmt.Errorf("role must be viewer, editor, or admin")
	}
	return s.users.UpdateRole(ctx, id, role)
}

func (s *AdminService) UpdateVehicle(
	ctx context.Context,
	id uuid.UUID,
	request dto.UpdateVehicleRequest,
) error {
	item, err := s.vehicles.GetByID(ctx, id)
	if err != nil {
		return err
	}

	item.VIN = request.VIN
	item.Year = request.Year
	item.Make = request.Make
	item.Model = request.Model
	item.Trim = request.Trim
	item.Engine = request.Engine
	item.Drivetrain = request.Drivetrain
	item.Transmission = request.Transmission
	item.Mileage = request.Mileage
	item.Color = request.Color
	item.LicensePlate = request.LicensePlate
	item.State = request.State

	return s.vehicles.Update(ctx, item)
}

func (s *AdminService) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	item, err := s.vehicles.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.vehicles.Delete(ctx, id, item.CustomerID)
}

func (s *AdminService) UpdateCustomer(
	ctx context.Context,
	id uuid.UUID,
	request dto.UpdateCustomerRequest,
) error {
	customers, err := s.customers.FindAll(ctx)
	if err != nil {
		return err
	}

	for index := range customers {
		if customers[index].ID != id {
			continue
		}
		item := &customers[index]
		item.FirstName = request.FirstName
		item.LastName = request.LastName
		item.Phone = request.Phone
		item.AddressLine1 = request.AddressLine1
		item.AddressLine2 = request.AddressLine2
		item.City = request.City
		item.State = request.State
		item.PostalCode = request.PostalCode
		return s.customers.Update(ctx, item)
	}

	return fmt.Errorf("customer not found")
}
