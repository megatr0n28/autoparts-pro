package repository

import (
	"context"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		u *user.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*user.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) UserRepository {

	return &userRepository{
		db: db,
	}

}

func (r *userRepository) Create(
	ctx context.Context,
	u *user.User,
) error {

	return r.db.
		WithContext(ctx).
		Create(u).
		Error

}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {

	var u user.User

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"email = ?",
				email,
			).
			First(&u).
			Error

	if err != nil {

		return nil, err

	}

	return &u, nil

}
