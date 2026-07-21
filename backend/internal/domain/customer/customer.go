package customer

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/domain/common"
)

type Customer struct {
	common.BaseEntity

	FirstName string `gorm:"size:100;not null"`

	LastName string `gorm:"size:100;not null"`

	Email string `gorm:"size:255;uniqueIndex"`

	Phone string `gorm:"size:30"`
}

func (Customer) TableName() string {

	return "customers"

}
