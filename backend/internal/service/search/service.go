package search

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	"github.com/megatr0n28/autoparts-pro/backend/internal/provider"
)

type Service struct {
	providers []provider.Provider
}

func New(
	providers ...provider.Provider,
) *Service {

	return &Service{
		providers: providers,
	}
}

func (s *Service) Search(
	ctx context.Context,
	vehicleID uuid.UUID,
	query string,
) ([]dto.PartSearchResponse, error) {

	var results []dto.PartSearchResponse

	for _, p := range s.providers {

		items, err :=
			p.Search(
				ctx,
				vehicleID,
				query,
			)

		if err != nil {
			continue
		}

		results =
			append(
				results,
				items...,
			)
	}

	return results, nil
}
