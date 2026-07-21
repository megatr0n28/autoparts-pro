package repository

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type postgresRepository[T any] struct {
	db *gorm.DB
}

func NewPostgresRepository[T any](
	db *gorm.DB,
) Repository[T] {

	return &postgresRepository[T]{
		db: db,
	}

}

func (r *postgresRepository[T]) Create(

	ctx context.Context,

	entity *T,

) error {

	return r.db.
		WithContext(ctx).
		Create(entity).
		Error

}

func (r *postgresRepository[T]) FindByID(

	ctx context.Context,

	id uuid.UUID,

	entity *T,

) error {

	return r.db.
		WithContext(ctx).
		First(
			entity,
			"id = ?",
			id,
		).
		Error

}

func (r *postgresRepository[T]) FindAll(

	ctx context.Context,

	entities *[]T,

) error {

	return r.db.
		WithContext(ctx).
		Find(
			entities,
		).
		Error

}

func (r *postgresRepository[T]) Update(

	ctx context.Context,

	entity *T,

) error {

	return r.db.
		WithContext(ctx).
		Save(entity).
		Error

}

func (r *postgresRepository[T]) Delete(

	ctx context.Context,

	entity *T,

) error {

	return r.db.
		WithContext(ctx).
		Delete(entity).
		Error

}
