package cloud

import (
	"github.com/softiron/manifold-api/deprecated/v2/cloud/datacenter"
)

// DataCenter is the API payload based on the legacy xmlrpc backend.
type DataCenter struct {
	ID           int                    `json:"id" yaml:"id"`
	Name         string                 `json:"name" yaml:"name"`
	Groups       []int                  `json:"groups" yaml:"groups"`
	Clusters     []datacenter.Cluster   `json:"clusters" yaml:"clusters"`
	Hosts        []datacenter.Host      `json:"hosts" yaml:"hosts"`
	Datastores   []datacenter.Datastore `json:"datastores" yaml:"datastores"`
	Networks     []datacenter.Network   `json:"networks" yaml:"networks"`
	Template     DataCenterTemplate     `json:"template" yaml:"template"`
	TemplateText string                 `json:"template_text" yaml:"template_text"`
}

// DataCenterTemplate is the API payload based on the legacy xmlrpc backend.
type DataCenterTemplate struct {
	Values      map[string]string `json:"values" yaml:"values"`
	Description string            `json:"description" yaml:"description"`
}
