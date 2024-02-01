// Package nettmpl provides strcuts for the VNetTemplate payload.
package nettmpl

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	VNetMAD string `json:"vnmad" yaml:"vnmad"`
}

// ParseTemplate return a structured Template based on the given map.
func ParseTemplate(m map[string]any) (*Template, error) {
	var t Template

	for key, value := range m {
		if v, ok := value.(string); ok {
			if key == "VN_MAD" {
				t.VNetMAD = v
			}
		}
	}

	return &t, nil
}
