package search

import (
	"strings"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
)

// dedupe removes duplicate parts based on PartNumber.
// If two providers return the same part number,
// only the first occurrence (already sorted by price)
// will be kept.
func dedupe(
	results []dto.PartSearchResponse,
) []dto.PartSearchResponse {

	seen := make(
		map[string]struct{},
	)

	unique := make(
		[]dto.PartSearchResponse,
		0,
		len(results),
	)

	for _, item := range results {

		key := strings.ToLower(
			strings.TrimSpace(
				item.PartNumber,
			),
		)

		// If part number is empty,
		// use retailer + name as a fallback.
		if key == "" {

			key =
				strings.ToLower(
					item.Retailer +
						":" +
						strings.TrimSpace(item.Name),
				)

		}

		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}

		unique = append(
			unique,
			item,
		)

	}

	return unique

}
