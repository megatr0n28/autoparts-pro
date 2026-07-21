package common

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func (entity *BaseEntity) BeforeCreate(
	tx *gorm.DB,
) error {

	if entity.ID == uuid.Nil {

		entity.ID = uuid.New()

	}

	return nil
}
