// Package insttmpl provides structs for the InstanceTemplate payload.
package insttmpl

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	Values             map[string]string `json:"values" yaml:"values"`
	VCenterCCRRef      string            `json:"vcenter_ccrref" yaml:"vcenter_ccrref"`
	VCenterInstanceID  string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterTemplateRef string            `json:"vcenter_template_ref" yaml:"vcenter_template_ref"`
}
