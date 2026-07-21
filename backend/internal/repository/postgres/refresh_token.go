package postgres

import (
	"context"

	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/refresh_token"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

type RefreshToken struct {
	ID uuid.UUID
}

func NewRefreshTokenRepository(
	db *gorm.DB,
) *RefreshTokenRepository {

	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(
	ctx context.Context,
	token *refresh_token.RefreshToken,
) error {

	return r.db.
		WithContext(ctx).
		Create(token).
		Error
}

func (r *RefreshTokenRepository) FindByHash(
	ctx context.Context,
	hash string,
) (*refresh_token.RefreshToken, error) {

	var token refresh_token.RefreshToken

	err :=
		r.db.
			WithContext(ctx).
			Where("token_hash = ?", hash).
			First(&token).
			Error

	return &token, err
}

func (r *RefreshTokenRepository) Revoke(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Model(&refresh_token.RefreshToken{}).
		Where("id = ?", id).
		Update("revoked", true).
		Error
}

func (r *RefreshTokenRepository) DeleteExpired(
	ctx context.Context,
) error {

	return r.db.
		WithContext(ctx).
		Where("expires_at < NOW()").
		Delete(&refresh_token.RefreshToken{}).
		Error
}

func (r *RefreshTokenRepository) DeleteUserTokens(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&refresh_token.RefreshToken{}).
		Error
}

func (r *RefreshToken) BeforeCreate(
	tx *gorm.DB,
) error {

	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	return nil
}

func (r *RefreshTokenRepository) FindActiveByUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]*refresh_token.RefreshToken, error) {

	var tokens []*refresh_token.RefreshToken

	err :=
		r.db.
			WithContext(ctx).
			Where(
				"user_id = ? AND revoked = false",
				userID,
			).
			Find(&tokens).
			Error

	return tokens, err

}
