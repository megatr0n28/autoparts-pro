package provider

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
)

type Provider interface {
	Name() string

	Search(
		ctx context.Context,
		vehicleID uuid.UUID,
		query string,
	) ([]dto.PartSearchResponse, error)
}
