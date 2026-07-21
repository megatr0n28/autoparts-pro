package refresh_token

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`

	TokenHash string `gorm:"not null;uniqueIndex"`

	ExpiresAt time.Time `gorm:"not null"`

	Revoked bool `gorm:"default:false"`

	DeviceName string

	IPAddress string

	CreatedAt time.Time

	UpdatedAt time.Time
}

// BeforeCreate generates UUID automatically
func (r *RefreshToken) BeforeCreate(
	tx *gorm.DB,
) error {

	if r.ID == uuid.Nil {

		r.ID = uuid.New()

	}

	return nil
}
