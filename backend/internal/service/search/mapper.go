package search

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	"github.com/megatr0n28/autoparts-pro/backend/internal/provider"
)

func mapProviderParts(
	parts []provider.Part,
) []dto.PartSearchResponse {

	results := make(
		[]dto.PartSearchResponse,
		0,
		len(parts),
	)

	for _, part := range parts {

		results = append(
			results,
			dto.PartSearchResponse{
				Retailer:    part.Retailer,
				Brand:       part.Brand,
				PartNumber:  part.PartNumber,
				Name:        part.Name,
				Description: part.Description,
				Price:       part.Price,
				Currency:    part.Currency,
				InStock:     part.InStock,
				ProductURL:  part.ProductURL,
				ImageURL:    part.ImageURL,
			},
		)
	}

	return results
}
