package cloud

import (
	"encoding/xml"

	"github.com/softiron/hypercloud-api/cloud/datacenter"
	"github.com/softiron/hypercloud-api/cloud/template"
)

// DataCenter is the API payload based on the legacy xmlrpc backend.
type DataCenter struct {
	ID         int                    `json:"id" yaml:"id"`
	Name       string                 `json:"name" yaml:"name"`
	Groups     []int                  `json:"groups" yaml:"groups"`
	Clusters   []datacenter.Cluster   `json:"clusters" yaml:"clusters"`
	Hosts      []datacenter.Host      `json:"hosts" yaml:"hosts"`
	Datastores []datacenter.Datastore `json:"datastores" yaml:"datastores"`
	Networks   []datacenter.Network   `json:"networks" yaml:"networks"`
	Template   DataCenterTemplate     `json:"template" yaml:"template"`
}

// DataCenterTemplate is the API payload based on the legacy xmlrpc backend.
type DataCenterTemplate struct {
	XMLName     xml.Name        `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values      template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty" xml:"DESCRIPTION,omitempty"`
}
