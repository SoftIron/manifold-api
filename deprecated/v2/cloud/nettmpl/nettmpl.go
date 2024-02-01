// Package nettmpl provides strcuts for the VNetTemplate payload.
package nettmpl

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	Values  map[string]string `json:"values" yaml:"values"`
	VNetMAD string            `json:"vnmad" yaml:"vnmad"`
}
