package vehicle

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	vehicleDomain "github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type Service struct {
	repo repository.VehicleRepository
}

func NewService(
	repo repository.VehicleRepository,
) *Service {

	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	v *vehicleDomain.Vehicle,
) error {

	return s.repo.Create(
		ctx,
		v,
	)
}

func (s *Service) List(
	ctx context.Context,
	customerID uuid.UUID,
) ([]vehicleDomain.Vehicle, error) {

	return s.repo.GetByCustomer(
		ctx,
		customerID,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	return s.repo.Delete(
		ctx,
		id,
		customerID,
	)
}

func (s *Service) SetPrimary(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	if err := s.repo.ClearPrimary(
		ctx,
		customerID,
	); err != nil {
		return err
	}

	return s.repo.SetPrimary(
		ctx,
		id,
		customerID,
	)
}

func (s *Service) Update(
	ctx context.Context,
	id uuid.UUID,
	customerID uuid.UUID,
	request vehicleDomain.Vehicle,
) error {

	existing, err :=
		s.repo.GetCustomerVehicle(
			ctx,
			id,
			customerID,
		)

	if err != nil {
		return err
	}

	if request.VIN != "" {

		request.VIN =
			vehicleDomain.NormalizeVIN(
				request.VIN,
			)

		found, err :=
			s.repo.FindByVIN(
				ctx,
				request.VIN,
			)

		if err == nil &&
			found.ID != existing.ID {

			return fmt.Errorf(
				"vehicle VIN already exists",
			)
		}
	}

	existing.VIN = request.VIN
	existing.Year = request.Year
	existing.Make = request.Make
	existing.Model = request.Model
	existing.Trim = request.Trim
	existing.Engine = request.Engine
	existing.Transmission = request.Transmission
	existing.Drivetrain = request.Drivetrain
	existing.Mileage = request.Mileage
	existing.Color = request.Color
	existing.LicensePlate = request.LicensePlate
	existing.State = request.State

	return s.repo.Update(
		ctx,
		existing,
	)
}
