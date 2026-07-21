package user

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/common"
)

type User struct {
	common.BaseEntity

	FirstName string `gorm:"size:100;not null"`

	LastName string `gorm:"size:100;not null"`

	Email string `gorm:"size:255;uniqueIndex;not null"`

	PasswordHash string `gorm:"size:255;not null"`

	Role string `gorm:"size:50;not null"`

	Active bool `gorm:"default:true"`
}

func (User) TableName() string {

	return "users"

}
