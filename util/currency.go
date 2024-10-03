package util

const (
	USD = "USD"
	RUB = "RUB"
	KZT = "KZT"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, RUB, KZT:
		return true
	}
	return false
}
