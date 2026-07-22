package vehicle

import "strings"

func NormalizeVIN(
	vin string,
) string {

	return strings.ToUpper(
		strings.TrimSpace(vin),
	)
}
