package service

import (
	"context"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type UserService struct {
	users repository.UserRepository
}

func NewUserService(
	users repository.UserRepository,
) *UserService {

	return &UserService{
		users: users,
	}

}

func (s *UserService) GetByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {

	return s.users.FindByEmail(
		ctx,
		email,
	)

}
