// Package hc documents the request and response payloads for version 1 of the sifi API.
package hc

// Root path for API endpoint.
const (
	PathPrefix            = "v1"
	AccessControlListPath = "acl"
	ClusterPath           = "cluster"
	DatastorePath         = "datastore"
	DocumentPath          = "document"
	GroupPath             = "group"
	HookPath              = "hook"
	HostPath              = "host"
	ImagePath             = "image"
	InstancePath          = "instance"
	InstanceGroupPath     = "instance-group"
	MarketPath            = "market"
	PingPath              = "ping"
	SecurityGroupPath     = "security-group"
	SystemPath            = "system"
	TemplatePath          = "template"
	UserPath              = "user"
	VirtualDataCenterPath = "vdc"
	VirtualNetworkPath    = "vnet"
	VirtualRouterPath     = "vrouter"
	ZonePath              = "zone"
)

// Period is a time interval with optional start and end times.
type Period struct { // TODO: use a time type for v2?
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// Perms is a set of owner (user), group, and other permissions. Think UNIX.
type Perms struct {
	OwnerUse    *bool `json:"owner_use,omitempty"`
	OwnerManage *bool `json:"owner_manage,omitempty"`
	OwnerAdmin  *bool `json:"owner_admin,omitempty"`
	GroupUse    *bool `json:"group_use,omitempty"`
	GroupManage *bool `json:"group_manage,omitempty"`
	GroupAdmin  *bool `json:"group_admin,omitempty"`
	OtherUse    *bool `json:"other_use,omitempty"`
	OtherManage *bool `json:"other_manage,omitempty"`
	OtherAdmin  *bool `json:"other_admin,omitempty"`
}
