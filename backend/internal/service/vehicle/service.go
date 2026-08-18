package vehicle

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/megatr0n28/autoparts-pro/backend/internal/authz"
	vehicleDomain "github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type Service struct {
	repo       repository.VehicleRepository
	authorizer *authz.Authorizer
}

func NewService(
	repo repository.VehicleRepository,
	authorizer *authz.Authorizer,
) *Service {
	return &Service{
		repo:       repo,
		authorizer: authorizer,
	}
}

// Create creates a vehicle and establishes the OpenFGA ownership
// relationship:
//
// user:<user-id> owner vehicle:<vehicle-id>
func (s *Service) Create(
	ctx context.Context,
	userID string,
	v *vehicleDomain.Vehicle,
) error {

	if v == nil {
		return fmt.Errorf("vehicle is required")
	}

	if userID == "" {
		return fmt.Errorf(
			"create vehicle authorization: authenticated user ID is missing",
		)
	}

	if s.authorizer == nil {
		return fmt.Errorf(
			"create vehicle authorization: authorizer is not configured",
		)
	}

	if v.VIN != "" {
		v.VIN = vehicleDomain.NormalizeVIN(v.VIN)

		found, err := s.repo.FindByVIN(ctx, v.VIN)

		if err == nil && found != nil {
			return fmt.Errorf("vehicle VIN already exists")
		}
	}

	if err := s.repo.Create(ctx, v); err != nil {
		return err
	}

	user := "user:" + userID
	object := "vehicle:" + v.ID.String()

	if err := s.authorizer.WriteTuple(
		ctx,
		user,
		"owner",
		object,
	); err != nil {
		return fmt.Errorf(
			"create vehicle authorization: %w",
			err,
		)
	}

	return nil
}

// List returns only vehicles that the authenticated user
// is authorized to access.
func (s *Service) List(
	ctx context.Context,
	userID string,
	customerID uuid.UUID,
) ([]vehicleDomain.Vehicle, error) {

	if userID == "" {
		return nil, fmt.Errorf(
			"list vehicles authorization: authenticated user ID is missing",
		)
	}

	if s.authorizer == nil {
		return nil, fmt.Errorf(
			"list vehicles authorization: authorizer is not configured",
		)
	}

	vehicles, err := s.repo.GetByCustomer(
		ctx,
		customerID,
	)

	if err != nil {
		return nil, err
	}

	user := "user:" + userID

	authorized := make(
		[]vehicleDomain.Vehicle,
		0,
		len(vehicles),
	)

	for _, vehicle := range vehicles {

		object := "vehicle:" + vehicle.ID.String()

		allowed, err := s.authorizer.Check(
			ctx,
			user,
			"owner",
			object,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"list vehicle authorization: %w",
				err,
			)
		}

		if allowed {
			authorized = append(
				authorized,
				vehicle,
			)
		}
	}

	return authorized, nil
}

// Delete deletes a vehicle only when the authenticated user
// owns that vehicle.
func (s *Service) Delete(
	ctx context.Context,
	userID string,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	if err := s.requireVehicleOwner(
		ctx,
		userID,
		id,
	); err != nil {
		return err
	}

	return s.repo.Delete(
		ctx,
		id,
		customerID,
	)
}

// SetPrimary sets a vehicle as primary only when the
// authenticated user owns the vehicle.
func (s *Service) SetPrimary(
	ctx context.Context,
	userID string,
	id uuid.UUID,
	customerID uuid.UUID,
) error {

	if err := s.requireVehicleOwner(
		ctx,
		userID,
		id,
	); err != nil {
		return err
	}

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

// Update updates a vehicle only when the authenticated user
// owns that vehicle.
func (s *Service) Update(
	ctx context.Context,
	userID string,
	id uuid.UUID,
	customerID uuid.UUID,
	request vehicleDomain.Vehicle,
) error {

	if err := s.requireVehicleOwner(
		ctx,
		userID,
		id,
	); err != nil {
		return err
	}

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
			found != nil &&
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

// requireVehicleOwner verifies:
//
// user:<userID> owner vehicle:<vehicleID>
func (s *Service) requireVehicleOwner(
	ctx context.Context,
	userID string,
	vehicleID uuid.UUID,
) error {

	if userID == "" {
		return fmt.Errorf(
			"vehicle authorization: authenticated user ID is missing",
		)
	}

	if s.authorizer == nil {
		return fmt.Errorf(
			"vehicle authorization: authorizer is not configured",
		)
	}

	user := "user:" + userID
	object := "vehicle:" + vehicleID.String()

	allowed, err := s.authorizer.Check(
		ctx,
		user,
		"owner",
		object,
	)

	if err != nil {
		return fmt.Errorf(
			"vehicle authorization: %w",
			err,
		)
	}

	if !allowed {
		return fmt.Errorf(
			"vehicle authorization: access denied",
		)
	}

	return nil
}
