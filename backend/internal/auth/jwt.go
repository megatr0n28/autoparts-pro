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

	claims := jwt.MapClaims{

		"user_id": userID,

		"role": role,

		"exp": time.Now().
			Add(j.Expiration).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(j.Secret),
	)

}
