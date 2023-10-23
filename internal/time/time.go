// Package time provides a wrapper for time.Time that implements the
// json.Marshaler interface. This fixes the problem where a zero time is
// marshaled as "0001-01-01T00:00:00Z" instead of respected the omitempty ta
//
// This doesn't fix omitempty, but will render the zero time as ""
package time

import (
	"time"
)

// Time is a wrapper for time.Time that implements the text.Marshaler interface.
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface for t.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	return t.Time.MarshalJSON()
}

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
}

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}
