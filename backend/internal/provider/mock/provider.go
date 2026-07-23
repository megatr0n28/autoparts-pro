package mock

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/provider"
)

type MockProvider struct{}

func New() *MockProvider {
	return &MockProvider{}
}

// Name returns the provider name.
func (m *MockProvider) Name() string {
	return "Mock Provider"
}

// Search returns mock search results.
// Later this implementation will be replaced by
// AutoZone, NAPA, Advance Auto, etc.
func (m *MockProvider) Search(
	ctx context.Context,
	vehicleID uuid.UUID,
	query string,
) ([]provider.Part, error) {

	// Honor context cancellation.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	query = strings.TrimSpace(strings.ToLower(query))

	// Example responses based on the search query.
	switch query {

	case "oil filter", "filter", "oil":

		return []provider.Part{
			{
				Retailer:    "AutoZone",
				Brand:       "FRAM",
				PartNumber:  "PH7317",
				Name:        "Engine Oil Filter",
				Description: "FRAM Extra Guard Spin-On Oil Filter",
				Price:       8.99,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/autozone/ph7317",
				ImageURL:    "",
			},
			{
				Retailer:    "Advance Auto",
				Brand:       "Mobil 1",
				PartNumber:  "M1-110A",
				Name:        "Extended Performance Oil Filter",
				Description: "Mobil 1 Extended Performance Filter",
				Price:       11.49,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/advance/m1-110a",
				ImageURL:    "",
			},
			{
				Retailer:    "NAPA",
				Brand:       "NAPA Gold",
				PartNumber:  "1348",
				Name:        "NAPA Gold Oil Filter",
				Description: "Premium engine oil filter",
				Price:       10.79,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/napa/1348",
				ImageURL:    "",
			},
		}, nil

	case "brake pads", "pads", "brake":

		return []provider.Part{
			{
				Retailer:    "AutoZone",
				Brand:       "Duralast",
				PartNumber:  "MKD905",
				Name:        "Ceramic Brake Pads",
				Description: "Front ceramic brake pad set",
				Price:       39.99,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/autozone/mkd905",
				ImageURL:    "",
			},
			{
				Retailer:    "Advance Auto",
				Brand:       "Carquest",
				PartNumber:  "CQ905",
				Name:        "Premium Ceramic Brake Pads",
				Description: "Front premium ceramic brake pads",
				Price:       42.49,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/advance/cq905",
				ImageURL:    "",
			},
		}, nil

	case "battery":

		return []provider.Part{
			{
				Retailer:    "AutoZone",
				Brand:       "Duralast Gold",
				PartNumber:  "H6-DLG",
				Name:        "Automotive Battery",
				Description: "650 CCA automotive battery",
				Price:       189.99,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/autozone/h6-dlg",
				ImageURL:    "",
			},
			{
				Retailer:    "NAPA",
				Brand:       "Legend Premium",
				PartNumber:  "BAT-7565",
				Name:        "Premium Battery",
				Description: "700 CCA maintenance-free battery",
				Price:       199.99,
				Currency:    "USD",
				InStock:     true,
				ProductURL:  "https://example.com/napa/bat7565",
				ImageURL:    "",
			},
		}, nil

	default:

		// Unknown search returns no results.
		return []provider.Part{}, nil
	}
}
