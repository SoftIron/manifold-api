//go:build !prod
// +build !prod

package snapshot

const isProduction = false

// IsValid returns true if p is a known time Period.
func (p Period) IsValid() bool {
	return p.IsAPeriod()
}
