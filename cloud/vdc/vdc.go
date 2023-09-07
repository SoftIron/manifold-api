// Package vdc provides structs for the VDC payload.
package vdc

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

// VNet is the API payload based on the legacy xmlrpc backend.
type VNet struct {
	Zone int `json:"zone" yaml:"zone"`
	VNet int `json:"vnet" yaml:"vnet"`
}
