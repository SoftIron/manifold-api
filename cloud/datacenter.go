package cloud

import (
	"github.com/softiron/manifold-api/cloud/datacenter"
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
	Template   Template               `json:"template" yaml:"template"`
}

// DataCenterTemplate is the API payload based on the legacy xmlrpc backend.
type DataCenterTemplate struct {
	Description string
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (d *DataCenter) ParseTemplate() (*DataCenterTemplate, error) {
	var t DataCenterTemplate

	for key, value := range d.Template {
		if v, ok := value.(string); ok {
			if key == "DESCRIPTION" {
				t.Description = v
			}
		}
	}

	return &t, nil
}
