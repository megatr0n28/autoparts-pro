package vehicle

import (
	"context"

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
