package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/refresh_token"
)

type RefreshTokenRepository interface {
	Create(
		ctx context.Context,
		token *refresh_token.RefreshToken,
	) error

	FindByHash(
		ctx context.Context,
		hash string,
	) (*refresh_token.RefreshToken, error)

	Revoke(
		ctx context.Context,
		id uuid.UUID,
	) error

	DeleteExpired(
		ctx context.Context,
	) error

	DeleteUserTokens(
		ctx context.Context,
		userID uuid.UUID,
	) error

	FindActiveByUser(
		ctx context.Context,
		userID uuid.UUID,
	) ([]*refresh_token.RefreshToken, error)
}
