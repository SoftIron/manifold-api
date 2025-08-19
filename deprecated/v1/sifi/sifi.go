// Package sifi implements the go language bindings for the Manifold API. The
// bindings will support exactly one version of the API at a time, even though
// the data types for all currently supported versions can be found in the v*
// directories.
//
// The generated directory contains the generated code from the legacy XMLRPC
// API. Once all of these types have been forked and moved into the models
// directory, this directory can be removed. This should be a prerequisite for
// moving this code to public github.
package sifi

import (
	"fmt"
	"strings"

	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

const (
	// PortNumber is the default port for `sifid` (sifi on a keypad).
	PortNumber = 7434
	// ServiceName is the public name of the SIFI service.
	ServiceName = "sifid"
)

// path returns a URL path with the correct prefix appended.
func path(dirs ...interface{}) string {
	p := make([]string, 1, len(dirs)+1)
	p[0] = hc.PathPrefix

	for _, d := range dirs {
		p = append(p, fmt.Sprint(d))
	}

	return strings.Join(p, "/")
}
