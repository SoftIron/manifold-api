package hc

// Cluster is the API payload based on the legacy xmlrpc backend.
type Cluster struct {
	ID         int               `json:"id" yaml:"id"`
	Name       string            `json:"name" yaml:"name"`
	Hosts      ClusterHosts      `json:"hosts" yaml:"hosts"`
	Datastores ClusterDatastores `json:"datastores" yaml:"datastores"`
	Vnets      ClusterVnets      `json:"vnets" yaml:"vnets"`
	Template   string            `json:"template" yaml:"template"`
}

// ClusterDatastores is the API payload based on the legacy xmlrpc backend.
type ClusterDatastores struct {
	ID []int `json:"id" yaml:"id"`
}

// ClusterHosts is the API payload based on the legacy xmlrpc backend.
type ClusterHosts struct {
	ID []int `json:"id" yaml:"id"`
}

// ClusterVnets is the API payload based on the legacy xmlrpc backend.
type ClusterVnets struct {
	ID []int `json:"id" yaml:"id"`
}
