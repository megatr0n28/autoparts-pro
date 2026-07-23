package mock

import (
	"context"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Name() string {
	return "Mock"
}

func (p *Provider) Search(
	ctx context.Context,
	vehicleID uuid.UUID,
	query string,
) ([]dto.PartSearchResponse, error) {

	return []dto.PartSearchResponse{

		{
			Retailer: "AutoZone",

			Brand: "FRAM",

			PartNumber: "PH7317",

			Name: "Engine Oil Filter",

			Description: "Premium oil filter",

			Price: 8.99,

			Currency: "USD",

			InStock: true,

			ProductURL: "https://example.com",

			ImageURL: "https://example.com/filter.jpg",
		},

		{
			Retailer: "Advance Auto",

			Brand: "Mobil 1",

			PartNumber: "M1-110A",

			Name: "Extended Performance Filter",

			Description: "Synthetic oil filter",

			Price: 11.49,

			Currency: "USD",

			InStock: true,

			ProductURL: "https://example.com",

			ImageURL: "https://example.com/filter2.jpg",
		},
	}, nil
}
