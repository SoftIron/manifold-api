// Package network provides structs for the VNet payload.
package network

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	Bridge         string
	BridgeType     string
	DNS            string
	Description    string
	Gateway        string
	Gateway6       string
	GuestMTU       int
	IP6Method      string
	IP6Metric      string
	Method         string
	Metric         string
	NetworkAddress string
	NetworkMask    string
	OuterVLANID    string
	SearchDomain   string
	SecurityGroups string
	VLAN           string
	VLANID         string
	VNMAD          string
}

// AddressRange is the API payload based on the legacy xmlrpc backend.
type AddressRange struct {
	ARID              string  `json:"arid" yaml:"arid"`
	GlobalPrefix      string  `json:"global_prefix" yaml:"global_prefix"`
	IP                string  `json:"ip" yaml:"ip"`
	MAC               string  `json:"mac" yaml:"mac"`
	ParentNetworkARID string  `json:"parent_network_arid" yaml:"parent_network_arid"`
	Size              int     `json:"size" yaml:"size"`
	Type              string  `json:"type" yaml:"type"`
	ULAPrefix         string  `json:"ula_prefix" yaml:"ula_prefix"`
	VNMAD             string  `json:"vnmad" yaml:"vnmad"`
	MACEnd            string  `json:"macend" yaml:"macend"`
	IPEnd             string  `json:"ipend" yaml:"ipend"`
	IP6ULA            string  `json:"ip6_ula" yaml:"ip6_ula"`
	IP6ALUEnd         string  `json:"ip6_ula_end" yaml:"ip6_ula_end"`
	IP6Global         string  `json:"ip6_global" yaml:"ip6_global"`
	IP6GlobalEnd      string  `json:"ip6_global_end" yaml:"ip6_global_end"`
	IP6               string  `json:"ip6" yaml:"ip6"`
	IP6End            string  `json:"ip6_end" yaml:"ip6_end"`
	PortStart         string  `json:"port_start" yaml:"port_start"`
	PortSize          string  `json:"port_size" yaml:"port_size"`
	UsedLeases        string  `json:"used_leases" yaml:"used_leases"`
	Leases            []Lease `json:"leases" yaml:"leases"`
}

// Lease is the API payload based on the legacy xmlrpc backend.
type Lease struct {
	IP        string `json:"ip" yaml:"ip"`
	IP6       string `json:"ip6" yaml:"ip6"`
	IP6Global string `json:"ip6_global" yaml:"ip6_global"`
	IP6Link   string `json:"ip6_link" yaml:"ip6_link"`
	IP6ULA    string `json:"ip6_ula" yaml:"ip6_ula"`
	MAC       string `json:"mac" yaml:"mac"`
	Instance  int    `json:"instance" yaml:"instance"`
	VNet      int    `json:"vnet" yaml:"vnet"`
	VRouter   int    `json:"vrouter" yaml:"vrouter"`
}
