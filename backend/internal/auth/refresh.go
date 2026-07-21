package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/refresh_token"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"

	"github.com/megatr0n28/autoparts-pro/backend/internal/security"
)

type RefreshTokenService struct {
	tokens repository.RefreshTokenRepository

	jwt *JWTManager

	expiration time.Duration
}

func NewRefreshTokenService(
	tokens repository.RefreshTokenRepository,
	jwt *JWTManager,
	expiration time.Duration,
) *RefreshTokenService {

	return &RefreshTokenService{

		tokens: tokens,

		jwt: jwt,

		expiration: expiration,
	}

}

func (s *RefreshTokenService) Create(
	ctx context.Context,
	userID uuid.UUID,
	device string,
	ip string,
) (string, error) {

	raw := make(
		[]byte,
		64,
	)

	_, err :=
		rand.Read(
			raw,
		)

	if err != nil {

		return "",
			err
	}

	token :=
		base64.URLEncoding.EncodeToString(
			raw,
		)

	refresh :=
		&refresh_token.RefreshToken{

			UserID: userID,

			TokenHash: security.HashToken(
				token,
			),

			ExpiresAt: time.Now().
				Add(
					s.expiration,
				),

			DeviceName: device,

			IPAddress: ip,
		}

	err =
		s.tokens.Create(
			ctx,
			refresh,
		)

	if err != nil {

		return "",
			err

	}

	return token, nil

}

func (s *RefreshTokenService) Validate(
	ctx context.Context,
	raw string,
) (*refresh_token.RefreshToken, error) {

	hash :=
		security.HashToken(
			raw,
		)

	token, err :=
		s.tokens.FindByHash(
			ctx,
			hash,
		)

	if err != nil {

		return nil, err

	}

	if token.Revoked {

		return nil,
			fmt.Errorf(
				"refresh token revoked",
			)

	}

	if time.Now().After(
		token.ExpiresAt,
	) {

		return nil,
			fmt.Errorf(
				"refresh token expired",
			)

	}

	return token, nil

}

func (s *RefreshTokenService) Rotate(
	ctx context.Context,
	old *refresh_token.RefreshToken,
	device string,
	ip string,
) (string, error) {

	err :=
		s.tokens.Revoke(
			ctx,
			old.ID,
		)

	if err != nil {

		return "",
			err

	}

	return s.Create(
		ctx,
		old.UserID,
		device,
		ip,
	)

}

func (s *RefreshTokenService) Revoke(
	ctx context.Context,
	id uuid.UUID,
) error {

	return s.tokens.Revoke(
		ctx,
		id,
	)

}

func (s *RefreshTokenService) DeleteUserTokens(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return s.tokens.DeleteUserTokens(
		ctx,
		userID,
	)

}
