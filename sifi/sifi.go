// Package sifi implements the go language bindings for the HyperCloud SIFI
// API. The bindings will support exactly one version of the API at a time, even
// though the data types for all currently supported versions can be found in
// the v* directories.
package sifi

const (
	APIPrefix   = "/hypercloud-api" // APIPrefix is the prefix for all API paths.
	APIVersion  = "v2"              // APIVersion is the prefix for all API paths (without the leading slash).
	PortNumber  = 7434              // PortNumber is the default port for `sifid` (sifi on a keypad).
	ServiceName = "sifid"           // ServiceName is the public name of the SIFI service.
)
