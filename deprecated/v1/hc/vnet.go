package hc

// VNetAR is the API payload based on the legacy xmlrpc backend.
type VNetAR struct {
	ARID              string     `json:"arid" yaml:"arid"`
	GlobalPrefix      string     `json:"global_prefix" yaml:"global_prefix"`
	IP                string     `json:"ip" yaml:"ip"`
	MAC               string     `json:"mac" yaml:"mac"`
	ParentNetworkARID string     `json:"parent_network_arid" yaml:"parent_network_arid"`
	Size              int        `json:"size" yaml:"size"`
	Type              string     `json:"type" yaml:"type"`
	UlaPrefix         string     `json:"ula_prefix" yaml:"ula_prefix"`
	VNMAD             string     `json:"vnmad" yaml:"vnmad"`
	MACEnd            string     `json:"macend" yaml:"macend"`
	IPEnd             string     `json:"ipend" yaml:"ipend"`
	IP6Ula            string     `json:"ip6_ula" yaml:"ip6_ula"`
	IP6UlaEnd         string     `json:"ip6_ula_end" yaml:"ip6_ula_end"`
	IP6Global         string     `json:"ip6_global" yaml:"ip6_global"`
	IP6GlobalEnd      string     `json:"ip6_global_end" yaml:"ip6_global_end"`
	IP6               string     `json:"ip6" yaml:"ip6"`
	IP6End            string     `json:"ip6_end" yaml:"ip6_end"`
	PortStart         string     `json:"port_start" yaml:"port_start"`
	PortSize          string     `json:"port_size" yaml:"port_size"`
	UsedLeases        string     `json:"used_leases" yaml:"used_leases"`
	Leases            VNetLeases `json:"leases" yaml:"leases"`
}

// VNetARPool is the API payload based on the legacy xmlrpc backend.
type VNetARPool struct {
	AR []VNetAR `json:"ar" yaml:"ar"`
}

// VNetClusters is the API payload based on the legacy xmlrpc backend.
type VNetClusters struct {
	ID []int `json:"id" yaml:"id"`
}

// VNetLease is the API payload based on the legacy xmlrpc backend.
type VNetLease struct {
	IP        string `json:"ip" yaml:"ip"`
	IP6       string `json:"ip6" yaml:"ip6"`
	IP6Global string `json:"ip6_global" yaml:"ip6_global"`
	IP6Link   string `json:"ip6_link" yaml:"ip6_link"`
	IP6Ula    string `json:"ip6_ula" yaml:"ip6_ula"`
	MAC       string `json:"mac" yaml:"mac"`
	VM        int    `json:"vm" yaml:"vm"`
	Vnet      int    `json:"vnet" yaml:"vnet"`
	Vrouter   int    `json:"vrouter" yaml:"vrouter"`
}

// VNetLeases is the API payload based on the legacy xmlrpc backend.
type VNetLeases struct {
	Lease []VNetLease `json:"lease" yaml:"lease"`
}

// VNetLock is the API payload based on the legacy xmlrpc backend.
type VNetLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// VNetPermissions is the API payload based on the legacy xmlrpc backend.
type VNetPermissions struct {
	OwnerU int `json:"owner_u" yaml:"owner_u"`
	OwnerM int `json:"owner_m" yaml:"owner_m"`
	OwnerA int `json:"owner_a" yaml:"owner_a"`
	GroupU int `json:"group_u" yaml:"group_u"`
	GroupM int `json:"group_m" yaml:"group_m"`
	GroupA int `json:"group_a" yaml:"group_a"`
	OtherU int `json:"other_u" yaml:"other_u"`
	OtherM int `json:"other_m" yaml:"other_m"`
	OtherA int `json:"other_a" yaml:"other_a"`
}

// VNetTemplate is the API payload based on the legacy xmlrpc backend.
type VNetTemplate struct {
	Values               map[string]string `json:"values" yaml:"values"`
	DNS                  string            `json:"dns" yaml:"dns"`
	Gateway              string            `json:"gateway" yaml:"gateway"`
	Gateway6             string            `json:"gateway6" yaml:"gateway6"`
	GuestMTU             int               `json:"guest_mtu" yaml:"guest_mtu"`
	IP6Method            string            `json:"ip6_method" yaml:"ip6_method"`
	IP6Metric            string            `json:"ip6_metric" yaml:"ip6_metric"`
	Method               string            `json:"method" yaml:"method"`
	Metric               string            `json:"metric" yaml:"metric"`
	NetworkAddress       string            `json:"network_address" yaml:"network_address"`
	NetworkMask          string            `json:"network_mask" yaml:"network_mask"`
	SearchDomain         string            `json:"search_domain" yaml:"search_domain"`
	VCenterFromWild      string            `json:"vcenter_from_wild" yaml:"vcenter_from_wild"`
	VCenterInstanceID    string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterNetRef        string            `json:"vcenter_net_ref" yaml:"vcenter_net_ref"`
	VCenterPortgroupType string            `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type"`
	VCenterTemplateRef   string            `json:"vcenter_template_ref" yaml:"vcenter_template_ref"`
}

// VNet is the API payload based on the legacy xmlrpc backend.
type VNet struct {
	ID                   int             `json:"id" yaml:"id"`
	UID                  int             `json:"uid" yaml:"uid"`
	GID                  int             `json:"gid" yaml:"gid"`
	Uname                string          `json:"uname" yaml:"uname"`
	Gname                string          `json:"gname" yaml:"gname"`
	Name                 string          `json:"name" yaml:"name"`
	Lock                 VNetLock        `json:"lock" yaml:"lock"`
	Permissions          VNetPermissions `json:"permissions" yaml:"permissions"`
	Clusters             VNetClusters    `json:"clusters" yaml:"clusters"`
	Bridge               string          `json:"bridge" yaml:"bridge"`
	BridgeType           string          `json:"bridge_type" yaml:"bridge_type"`
	State                int             `json:"state" yaml:"state"`
	PrevState            int             `json:"prev_state" yaml:"prev_state"`
	ParentNetworkID      string          `json:"parent_network_id" yaml:"parent_network_id"`
	VNMAD                string          `json:"vnmad" yaml:"vnmad"`
	Phydev               string          `json:"phydev" yaml:"phydev"`
	VLANID               string          `json:"vlanid" yaml:"vlanid"`
	OuterVLANID          string          `json:"outer_vlanid" yaml:"outer_vlanid"`
	VLANIDAutomatic      string          `json:"vlanidautomatic" yaml:"vlanidautomatic"`
	OuterVLANIDAutomatic string          `json:"outer_vlanidautomatic" yaml:"outer_vlanidautomatic"`
	UsedLeases           int             `json:"used_leases" yaml:"used_leases"`
	Vrouters             VNetVRouters    `json:"vrouters" yaml:"vrouters"`
	Template             VNetTemplate    `json:"template" yaml:"template"`
	ARPool               VNetARPool      `json:"arpool" yaml:"arpool"`
}

// VNetVRouters is the API payload based on the legacy xmlrpc backend OneVrouters.
type VNetVRouters struct {
	ID []int `json:"id" yaml:"id"`
}
