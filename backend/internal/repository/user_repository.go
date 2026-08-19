package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (*user.User, error)

	Create(
		ctx context.Context,
		u *user.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*user.User, error)

	FindAll(
		ctx context.Context,
	) ([]user.User, error)

	UpdateRole(
		ctx context.Context,
		id uuid.UUID,
		role string,
	) error
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

func (r *userRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
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

func (r *userRepository) FindAll(
	ctx context.Context,
) ([]user.User, error) {
	var users []user.User
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateRole(
	ctx context.Context,
	id uuid.UUID,
	role string,
) error {
	return r.db.WithContext(ctx).
		Model(&user.User{}).
		Where("id = ?", id).
		Update("role", role).Error
}
