// Package api implements common internal types across the Manifold API.
package api

import (
	"fmt"
	"strconv"
	"strings"
)

// Str2Bool converts a string to a bool extending strconv.ParseBool to
// include "yes"/"no". It returns an error if the string is not a
// valid bool.
func Str2Bool(s string) (bool, error) {
	if s == "" || strings.EqualFold(s, "no") {
		return false, nil
	}

	if strings.EqualFold(s, "yes") {
		return true, nil
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("str2bool: %w", err)
	}

	return b, nil
}
