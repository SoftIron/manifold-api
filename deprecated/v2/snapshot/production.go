//go:build prod
// +build prod

package snapshot

const isProduction = true

// IsValid returns true if p is a known time Period.
func (p Period) IsValid() bool {
	return p >= Hourly && p.IsAPeriod() // disable Minutely
}
