// Package vnet provides structs for the VNet payload.
package vnet

import (
	"encoding/xml"

	"github.com/softiron/hypercloud-api/cloud/template"
)

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	XMLName              xml.Name        `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values               template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	Bridge               string          `json:"bridge" yaml:"bridge" xml:"BRIDGE"`
	BridgeType           string          `json:"bridge_type" yaml:"bridge_type" xml:"BRIDGE_TYPE"`
	DNS                  string          `json:"dns" yaml:"dns" xml:"DNS"`
	Description          string          `json:"description" yaml:"description" xml:"DESCRIPTION"`
	Gateway              string          `json:"gateway" yaml:"gateway" xml:"GATEWAY"`
	Gateway6             string          `json:"gateway6" yaml:"gateway6" xml:"GATEWAY6"`
	GuestMTU             int             `json:"guest_mtu" yaml:"guest_mtu" xml:"GUEST_MTU"`
	IP6Method            string          `json:"ip6_method" yaml:"ip6_method" xml:"IP6_METHOD"`
	IP6Metric            string          `json:"ip6_metric" yaml:"ip6_metric" xml:"IP6_METRIC"`
	Method               string          `json:"method" yaml:"method" xml:"METHOD"`
	Metric               string          `json:"metric" yaml:"metric" xml:"METRIC"`
	NetworkAddress       string          `json:"network_address" yaml:"network_address" xml:"NETWORK_ADDRESS"`
	NetworkMask          string          `json:"network_mask" yaml:"network_mask" xml:"NETWORK_MASK"`
	OuterVLANID          string          `json:"outer_vlan_id" yaml:"outer_vlan_id" xml:"OUTER_VLAN_ID"`
	SearchDomain         string          `json:"search_domain" yaml:"search_domain" xml:"SEARCH_DOMAIN"`
	SecurityGroups       string          `json:"security_groups" yaml:"security_groups" xml:"SECURITY_GROUPS"`
	VCenterFromWild      string          `json:"vcenter_from_wild" yaml:"vcenter_from_wild" xml:"VCENTER_FROM_WILD,omitempty"`
	VCenterInstanceID    string          `json:"vcenter_instance_id" yaml:"vcenter_instance_id" xml:"VCENTER_INSTANCE_ID,omitempty"`
	VCenterNetRef        string          `json:"vcenter_net_ref" yaml:"vcenter_net_ref" xml:"VCENTER_NET_REF,omitempty"`
	VCenterPortgroupType string          `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type" xml:"VCENTER_PORTGROUP_TYPE,omitempty"`
	VCenterTemplateRef   string          `json:"vcenter_template_ref" yaml:"vcenter_template_ref" xml:"VCENTER_TEMPLATE_REF,omitempty"`
	VLAN                 string          `json:"vlan" yaml:"vlan" xml:"VLAN"`
	VLANID               string          `json:"vlan_id" yaml:"vlan_id" xml:"VLAN_ID"`
	VNMAD                string          `json:"vn_mad" yaml:"vn_mad" xml:"VN_MAD"`
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
