package util

// Constants for all supported  currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	// In case currency is USD, EUR, CAD, we'll return true
	case USD, EUR, CAD:
		return true
	}
	return false
}
