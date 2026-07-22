package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/customer"
	"github.com/megatr0n28/autoparts-pro/backend/internal/security"
)

type Service struct {
	users repository.UserRepository

	jwt       *JWTManager
	refresh   *RefreshTokenService
	customers repository.CustomerRepository
}

func NewService(
	users repository.UserRepository,
	jwt *JWTManager,
	refresh *RefreshTokenService,
	customers repository.CustomerRepository,
) *Service {

	return &Service{
		users:     users,
		jwt:       jwt,
		refresh:   refresh,
		customers: customers,
	}

}

func (s *Service) Register(
	ctx context.Context,
	u *user.User,
	password string,
) error {

	hash, err :=
		security.HashPassword(
			password,
		)

	if err != nil {
		return err
	}

	u.PasswordHash = hash

	if u.Role == "" {
		u.Role = "user"
	}

	err =
		s.users.Create(
			ctx,
			u,
		)

	if err != nil {

		return err

	}

	profile :=
		&customer.Customer{

			UserID: u.ID,

			FirstName: u.FirstName,

			LastName: u.LastName,

			Country: "USA",
		}

	return s.customers.Create(
		ctx,
		profile,
	)

}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
	device string,
	ip string,
) (string, string, error) {

	u, err :=
		s.users.FindByEmail(
			ctx,
			email,
		)

	if err != nil {
		return "", "", err
	}

	err =
		security.ComparePassword(
			u.PasswordHash,
			password,
		)

	if err != nil {
		return "", "", err
	}

	accessToken, err :=
		s.jwt.GenerateToken(
			u.ID.String(),
			u.Role,
		)

	if err != nil {
		return "", "", err
	}

	refreshToken, err :=
		s.refresh.Create(
			ctx,
			u.ID,
			device,
			ip,
		)

	if err != nil {

		return "",
			"",
			err
	}

	return accessToken, refreshToken, nil
}

func (s *Service) Refresh(
	ctx context.Context,
	raw string,
	device string,
	ip string,
) (string, string, error) {

	old,
		err :=
		s.refresh.Validate(
			ctx,
			raw,
		)

	if err != nil {

		return "",
			"",
			err

	}

	newRefresh,
		err :=
		s.refresh.Rotate(
			ctx,
			old,
			device,
			ip,
		)

	if err != nil {

		return "",
			"",
			err

	}

	access,
		err :=
		s.jwt.GenerateToken(
			old.UserID.String(),
			"user",
		)

	if err != nil {

		return "",
			"",
			err

	}

	return access, newRefresh, nil

}

func (s *Service) Logout(
	ctx context.Context,
	raw string,
) error {

	token,
		err :=
		s.refresh.Validate(
			ctx,
			raw,
		)

	if err != nil {

		return err

	}

	return s.refresh.Revoke(
		ctx,
		token.ID,
	)

}

func (s *Service) LogoutAll(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return s.refresh.DeleteUserTokens(
		ctx,
		userID,
	)

}
