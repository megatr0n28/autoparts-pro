package auth

import (
	"context"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"

	"github.com/megatr0n28/autoparts-pro/backend/internal/security"
)

type Service struct {
	users repository.UserRepository

	jwt *JWTManager
}

func NewService(
	users repository.UserRepository,
	jwt *JWTManager,
) *Service {

	return &Service{
		users: users,
		jwt:   jwt,
	}
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	u, err :=
		s.users.FindByEmail(
			ctx,
			email,
		)

	if err != nil {
		return "", err
	}

	err =
		security.ComparePassword(
			u.PasswordHash,
			password,
		)

	if err != nil {
		return "", err
	}

	return s.jwt.GenerateToken(
		u.ID.String(),
		u.Role,
	)

}
