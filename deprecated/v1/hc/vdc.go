package hc

// ClusterVDC is the API payload based on the legacy xmlrpc backend.
type ClusterVDC struct {
	ZoneID    int `json:"zone_id" yaml:"zone_id"`
	ClusterID int `json:"cluster_id" yaml:"cluster_id"`
}

// VDCClusters is the API payload based on the legacy xmlrpc backend.
type VDCClusters struct {
	Cluster []ClusterVDC `json:"cluster" yaml:"cluster"`
}

// VDCDatastore is the API payload based on the legacy xmlrpc backend.
type VDCDatastore struct {
	ZoneID      int `json:"zone_id" yaml:"zone_id"`
	DatastoreID int `json:"datastore_id" yaml:"datastore_id"`
}

// VDCDatastores is the API payload based on the legacy xmlrpc backend.
type VDCDatastores struct {
	Datastore []VDCDatastore `json:"datastore" yaml:"datastore"`
}

// VDCGroups is the API payload based on the legacy xmlrpc backend.
type VDCGroups struct {
	ID []int `json:"id" yaml:"id"`
}

// VDCHost is the API payload based on the legacy xmlrpc backend.
type VDCHost struct {
	ZoneID int `json:"zone_id" yaml:"zone_id"`
	HostID int `json:"host_id" yaml:"host_id"`
}

// VDCHosts is the API payload based on the legacy xmlrpc backend.
type VDCHosts struct {
	Host []VDCHost `json:"host" yaml:"host"`
}

// VDC is the API payload based on the legacy xmlrpc backend.
type VDC struct {
	ID         int           `json:"id" yaml:"id"`
	Name       string        `json:"name" yaml:"name"`
	Groups     VDCGroups     `json:"groups" yaml:"groups"`
	Clusters   VDCClusters   `json:"clusters" yaml:"clusters"`
	Hosts      VDCHosts      `json:"hosts" yaml:"hosts"`
	Datastores VDCDatastores `json:"datastores" yaml:"datastores"`
	Vnets      VDCVnets      `json:"vnets" yaml:"vnets"`
	Template   string        `json:"template" yaml:"template"`
}

// VDCVnet is the API payload based on the legacy xmlrpc backend.
type VDCVnet struct {
	ZoneID int `json:"zone_id" yaml:"zone_id"`
	VnetID int `json:"vnet_id" yaml:"vnet_id"`
}

// VDCVnets is the API payload based on the legacy xmlrpc backend.
type VDCVnets struct {
	Vnet []VDCVnet `json:"vnet" yaml:"vnet"`
}
