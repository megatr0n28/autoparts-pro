package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	Secret string

	Expiration time.Duration
}

func NewJWTManager(
	secret string,
	expiration time.Duration,
) *JWTManager {

	return &JWTManager{
		Secret:     secret,
		Expiration: expiration,
	}

}

func (j *JWTManager) GenerateToken(
	userID string,
	role string,
) (string, error) {

	claims := Claims{

		UserID: userID,

		Role: role,

		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(
				time.Now().
					Add(j.Expiration),
			),

			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token :=
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			claims,
		)

	return token.SignedString(
		[]byte(j.Secret),
	)

}

func (j *JWTManager) ValidateToken(
	tokenString string,
) (*Claims, error) {

	token, err :=
		jwt.ParseWithClaims(
			tokenString,
			&Claims{},
			func(token *jwt.Token) (
				interface{},
				error,
			) {

				return []byte(j.Secret), nil

			},
		)

	if err != nil {
		return nil, err
	}

	claims, ok :=
		token.Claims.(*Claims)

	if !ok || !token.Valid {

		return nil,
			jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
