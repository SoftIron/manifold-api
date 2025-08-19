// Package manifold implements the Go language bindings for the Manifold API. The
// bindings will support exactly one version of the API at a time, even though
// the data types for all currently supported versions can be found in the v*
// directories.
package manifold

const (
	// APIPrefix is the prefix for all API paths.
	APIPrefix = "/manifold-api"
	// APIVersion is the prefix for all API paths (without the leading slash).
	APIVersion = "v3-preview"
	// PortNumber is the default port for `sifid` (sifi on a keypad).
	PortNumber = 7434
	// ServiceName is the public name of the manifold-api service.
	ServiceName = "sifid"
)
