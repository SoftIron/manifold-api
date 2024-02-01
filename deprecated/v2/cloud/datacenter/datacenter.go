// Package datacenter provides structs for the Datacenter payload.
package datacenter

// Cluster is the API payload based on the legacy xmlrpc backend.
type Cluster struct {
	Zone    int `json:"zone" yaml:"zone"`
	Cluster int `json:"cluster" yaml:"cluster"`
}

// Datastore is the API payload based on the legacy xmlrpc backend.
type Datastore struct {
	Zone      int `json:"zone" yaml:"zone"`
	Datastore int `json:"datastore" yaml:"datastore"`
}

// Host is the API payload based on the legacy xmlrpc backend.
type Host struct {
	Zone int `json:"zone" yaml:"zone"`
	Host int `json:"host" yaml:"host"`
}

// Network is the API payload based on the legacy xmlrpc backend.
type Network struct {
	Zone    int `json:"zone" yaml:"zone"`
	Network int `json:"network" yaml:"network"`
}
