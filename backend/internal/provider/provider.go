package provider

import (
	"context"

	"github.com/google/uuid"
)

// Provider represents a searchable auto parts retailer.
//
// Every retailer implementation (Mock, AutoZone,
// NAPA, Advance Auto, O'Reilly, etc.) must satisfy
// this interface.
type Provider interface {

	// Name returns the retailer name.
	Name() string

	// Search searches the retailer catalog.
	Search(
		ctx context.Context,
		vehicleID uuid.UUID,
		query string,
	) ([]Part, error)
}
