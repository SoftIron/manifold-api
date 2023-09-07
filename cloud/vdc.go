package cloud

import (
	"encoding/xml"

	"github.com/softiron/hypercloud-api/cloud/template"
	"github.com/softiron/hypercloud-api/cloud/vdc"
)

// VDC is the API payload based on the legacy xmlrpc backend.
type VDC struct {
	ID         int             `json:"id" yaml:"id"`
	Name       string          `json:"name" yaml:"name"`
	Groups     []int           `json:"groups" yaml:"groups"`
	Clusters   []vdc.Cluster   `json:"clusters" yaml:"clusters"`
	Hosts      []vdc.Host      `json:"hosts" yaml:"hosts"`
	Datastores []vdc.Datastore `json:"datastores" yaml:"datastores"`
	VNets      []vdc.VNet      `json:"vnets" yaml:"vnets"`
	Template   VDCTemplate     `json:"template" yaml:"template"`
}

// VDCTemplate is the API payload based on the legacy xmlrpc backend.
type VDCTemplate struct {
	XMLName     xml.Name        `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values      template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty" xml:"DESCRIPTION,omitempty"`
}
