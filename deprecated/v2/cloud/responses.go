package cloud

import (
	"github.com/softiron/hypercloud-api/deprecated/v2/cloud/instance"
	"github.com/softiron/hypercloud-api/deprecated/v2/cloud/nettmpl"
	"github.com/softiron/hypercloud-api/internal/time"
)

//
// ACL
//

// CreateACLResponse is the response body for POST /cloud/acl.
type CreateACLResponse struct {
	ACL int `json:"acl"`
}

// ACLsResponse is the response body for GET /cloud/acl.
type ACLsResponse struct {
	ACLs []ACL `json:"acls"`
}

//
// Cluster
//

// ClusterResponse is the response body for GET /cloud/cluster.
type ClusterResponse struct {
	Cluster Cluster `json:"cluster"`
}

// ClustersResponse is the response body for GET /cloud/cluster.
type ClustersResponse struct {
	Clusters []Cluster `json:"cluster"`
}

// CreateClusterResponse is response body for POST /cloud/cluster.
type CreateClusterResponse struct {
	Cluster int `json:"cluster"`
}

// UpdateClusterResponse is response body for PATCH /cloud/cluster.
type UpdateClusterResponse struct {
	Cluster int `json:"cluster"`
}

// AddClusterHostResponse is the response body for PATCH /cloud/cluster/host.
type AddClusterHostResponse struct {
	Host int `json:"host"`
}

// AddClusterDatastoreResponse is the response body for PATCH /cloud/cluster/datastore.
type AddClusterDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// AddClusterNetworkResponse is the response body for PATCH /cloud/cluster/vnet.
type AddClusterNetworkResponse struct {
	VNet int `json:"vnet"`
}

// RenameClusterResponse is the response body for PATCH /cloud/cluster/name.
type RenameClusterResponse struct {
	Cluster int `json:"cluster"`
}

//
// Datastore
//

// DatastoreResponse is the response body for GET /cloud/datastore.
type DatastoreResponse struct {
	Datastore Datastore `json:"datastore"`
}

// DatastoresResponse is the response body for GET /cloud/datastore.
type DatastoresResponse struct {
	Datastores []Datastore `json:"datastore"`
}

// CreateDatastoreResponse is the response body for POST /cloud/datastore.
type CreateDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// UpdateDatastoreResponse is the response body for PATCH /cloud/datastore.
type UpdateDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// ChangeDatastorePermissionsResponse is the response body for PATCH /cloud/datastore/permissions.
type ChangeDatastorePermissionsResponse struct {
	Datastore int `json:"datastore"`
}

// ChangeDatastoreOwnershipResponse is the response body for PATCH /cloud/datastore/ownership.
type ChangeDatastoreOwnershipResponse struct {
	Datastore int `json:"datastore"`
}

// RenameDatastoreResponse is the response body for PATCH /cloud/datastore/name.
type RenameDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// EnableDatastoreResponse is the response body for PATCH /cloud/datastore/enable.
type EnableDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

//
// Document
//

// DocumentResponse is the response body for GET /cloud/document.
type DocumentResponse struct {
	Document Document `json:"document"`
}

// DocumentsResponse is the response body for GET /cloud/document.
type DocumentsResponse struct {
	Documents []Document `json:"document"`
}

// AllocateDocumentResponse is the response body for POST /cloud/document.
type AllocateDocumentResponse struct {
	Document int `json:"document"`
}

// CloneDocumentResponse is the response body for POST /cloud/document/clone.
type CloneDocumentResponse struct {
	Document int `json:"document"`
}

// UpdateDocumentResponse is the response body for PATCH /cloud/document.
type UpdateDocumentResponse struct {
	Document int `json:"document"`
}

// ChangeDocumentPermissionsResponse is the response body for PATCH /cloud/document/permissions.
type ChangeDocumentPermissionsResponse struct {
	Document int `json:"document"`
}

// ChangeDocumentOwnershipResponse is the response body for PATCH /cloud/document/ownership.
type ChangeDocumentOwnershipResponse struct {
	Document int `json:"document"`
}

// RenameDocumentResponse is the response body for PATCH /cloud/document/name.
type RenameDocumentResponse struct {
	Document int `json:"document"`
}

// LockDocumentResponse is the response body for PATCH /cloud/document/lock.
type LockDocumentResponse struct {
	Document int       `json:"document"`
	Time     time.Time `json:"time"`
}

// UnlockDocumentResponse is the response body for PATCH /cloud/document/unlock.
type UnlockDocumentResponse struct {
	Document int `json:"document"`
}

//
// Group
//

// GroupResponse is the response body for GET /cloud/group.
type GroupResponse struct {
	Group Group `json:"group"`
}

// GroupsResponse is the response body for GET /cloud/group.
type GroupsResponse struct {
	Groups []Group `json:"groups"`
}

// GroupQuotaResponse is the response body for GET /cloud/group/quota.
type GroupQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

// CreateGroupResponse is the response body for POST /cloud/group.
type CreateGroupResponse struct {
	Group int `json:"group"`
}

// UpdateGroupResponse is the response body for PATCH /cloud/group.
type UpdateGroupResponse struct {
	Group int `json:"group"`
}

// UpdateGroupQuotaResponse is the response body for PATCH /cloud/group/{group}/quota.
type UpdateGroupQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quotas"`
}

// AddGroupAdminResponse is the response body for POST /cloud/group/admin.
type AddGroupAdminResponse struct {
	Group int `json:"group"`
}

// SetGroupQuotaResponse is the response body for POST /cloud/group/quota.
type SetGroupQuotaResponse struct {
	Group int `json:"group"`
}

//
// Hook
//

// HookResponse is the response body for GET /cloud/hook.
type HookResponse struct {
	Hook Hook `json:"hook"`
}

// HooksResponse is the response body for GET /cloud/hook.
type HooksResponse struct {
	Hooks []Hook `json:"hook"`
}

// HookLogResponse is the response body for GET /cloud/hook/log.
type HookLogResponse struct {
	HookLogs []HookLog `json:"hook_log"`
}

// CreateHookResponse is the response body for POST /cloud/hook.
type CreateHookResponse struct {
	Hook int `json:"hook"`
}

// UpdateHookResponse is the response body for PATCH /cloud/hook.
type UpdateHookResponse struct {
	Hook int `json:"hook"`
}

// RenameHookResponse is the response body for PATCH /cloud/hook/name.
type RenameHookResponse struct {
	Hook int `json:"hook"`
}

// LockHookResponse is the response body for PATCH /cloud/hook/lock.
type LockHookResponse struct {
	Hook int       `json:"hook"`
	Time time.Time `json:"time"`
}

// UnlockHookResponse is the response body for PATCH /cloud/hook/unlock.
type UnlockHookResponse struct {
	Hook int `json:"hook"`
}

// RetryHookResponse is the response body for PATCH /cloud/hook/retry.
type RetryHookResponse struct {
	Hook int `json:"hook"`
}

//
// Compute (v1 called this Host, but this is just the compute resources)
//

// HostResponse is the response body for GET /cloud/host.
type HostResponse struct {
	Host Host `json:"host"`
}

// HostsResponse is the response body for GET /cloud/host.
type HostsResponse struct {
	Hosts []Host `json:"hosts"`
}

// HostMonitoringResponse is the response body for GET /cloud/host/monitoring.
type HostMonitoringResponse struct {
	Monitoring []HostMonitoring `json:"monitoring"`
}

// HostsMonitoringResponse is the response body for GET /cloud/host/monitoring.
type HostsMonitoringResponse struct {
	Monitoring []HostMonitoring `json:"monitoring"`
}

// CreateHostResponse is the response body for POST /cloud/host.
type CreateHostResponse struct {
	Host int `json:"host"`
}

// SetHostStatusResponse is the response body for PATCH /cloud/host/status.
type SetHostStatusResponse struct {
	Host int `json:"host"`
}

// UpdateHostResponse is the response body for PATCH /cloud/host.
type UpdateHostResponse struct {
	Host int `json:"host"`
}

// RenameHostResponse is the response body for PATCH /cloud/host/name.
type RenameHostResponse struct {
	Host int `json:"host"`
}

//
// Image
//

// ImageResponse is the response body for GET /cloud/image.
type ImageResponse struct {
	Image Image `json:"image"`
}

// ImagesResponse is the response body for GET /cloud/image.
type ImagesResponse struct {
	Images []Image `json:"images"`
}

// CreateImageResponse is the response body for POST /cloud/image.
type CreateImageResponse struct {
	Image int `json:"image"`
}

// CloneImageResponse is the response body for POST /cloud/image/clone.
type CloneImageResponse struct {
	Image int `json:"image"`
}

// EnableImageResponse is the response body for PATCH /cloud/image/enable.
type EnableImageResponse struct {
	Image int `json:"image"`
}

// SetImagePersistentResponse is the response body for PATCH /cloud/image/persistent.
type SetImagePersistentResponse struct {
	Image int `json:"image"`
}

// ChangeImageTypeResponse is the response body for PATCH /cloud/image/type.
type ChangeImageTypeResponse struct {
	Image int `json:"image"`
}

// UpdateImageResponse is the response body for PATCH /cloud/image.
type UpdateImageResponse struct {
	Image int `json:"image"`
}

// ChangeImagePermissionsResponse is the response body for PATCH /cloud/image/permissions.
type ChangeImagePermissionsResponse struct {
	Image int `json:"image"`
}

// ChangeImageOwnershipResponse is the response body for PATCH /cloud/image/ownership.
type ChangeImageOwnershipResponse struct {
	Image int `json:"image"`
}

// RenameImageResponse is the response body for PATCH /cloud/image/name.
type RenameImageResponse struct {
	Image int `json:"image"`
}

// RevertImageSnapshotResponse is the response body for PATCH /cloud/image/snapshot/revert.
type RevertImageSnapshotResponse struct {
	Image int `json:"image"`
}

// FlattenImageSnapshotResponse is the response body for PATCH /cloud/image/snapshot/flatten.
type FlattenImageSnapshotResponse struct {
	Image int `json:"image"`
}

// LockImageResponse is the response body for PATCH /cloud/image/lock.
type LockImageResponse struct {
	Image int       `json:"image"`
	Time  time.Time `json:"time"`
}

// UnlockImageResponse is the response body for PATCH /cloud/image/unlock.
type UnlockImageResponse struct {
	Image int `json:"image"`
}

//
// Instance
//

// InstancesResponse is the response body for GET /cloud/instance.
type InstancesResponse struct {
	Instances []LockedInstance `json:"instances"`
}

// CreateInstanceResponse is the response body for POST /cloud/instance.
type CreateInstanceResponse struct {
	Instance int `json:"instance"`
}

// InstancesAccountingResponse is the response body for GET /cloud/instance/accounting.
type InstancesAccountingResponse struct {
	Accounting []AcctHistory `json:"accounting"`
}

// SetInstanceActionResponse is the response body for POST /cloud/instance/action.
type SetInstanceActionResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceConfigResponse is the response body for PATCH /cloud/instance/config.
type UpdateInstanceConfigResponse struct {
	Instance int `json:"instance"`
}

// DeployInstanceResponse is the response body for PATCH /cloud/instance/deploy.
type DeployInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskResponse is the response body for POST /cloud/instance/disk.
type CreateInstanceDiskResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskImageResponse is the response body for POST /cloud/instance/disk/image.
type CreateInstanceDiskImageResponse struct {
	Image int `json:"image"`
}

// ResizeInstanceDiskResponse is the response body for PATCH /cloud/instance/disk/size.
type ResizeInstanceDiskResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskSnapshotResponse is the response body for POST /cloud/instance/disk/snapshot.
type CreateInstanceDiskSnapshotResponse struct {
	Instance int `json:"instance"`
}

// RenameInstanceDiskSnapshotResponse is the response body for PATCH /cloud/instance/disk/snapshot/name.
type RenameInstanceDiskSnapshotResponse struct {
	Instance int `json:"instance"`
}

// RevertInstanceDiskSnapshotResponse is the response body for PATCH /cloud/instance/disk/snapshot/revert.
type RevertInstanceDiskSnapshotResponse struct {
	Snapshot int `json:"snapshot"`
}

// LockInstanceResponse is the response body for PATCH /cloud/instance/lock.
type LockInstanceResponse struct {
	Instance int       `json:"instance"`
	Time     time.Time `json:"time"`
}

// InstancesMonitoringResponse is the response body for GET /cloud/instance/monitoring.
type InstancesMonitoringResponse struct {
	Monitoring []instance.Monitoring `json:"monitoring"`
}

// InstanceMonitoringResponse is the response body for GET /cloud/instance/monitoring/{id}.
type InstanceMonitoringResponse struct {
	Monitoring []instance.Monitoring `json:"monitoring"`
}

// MoveInstanceResponse is the response body for PATCH /cloud/instance/move.
type MoveInstanceResponse struct {
	Instance int `json:"instance"`
}

// RenameInstanceResponse is the response body for PATCH /cloud/instance/name.
type RenameInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceNICResponse is the response body for POST /cloud/instance/nic.
type CreateInstanceNICResponse struct {
	Instance int `json:"instance"` // TODO: would make more sense as NIC ID, maybe docs are wrong
}

// ChangeInstanceOwnershipResponse is the response body for PATCH /cloud/instance/ownership.
type ChangeInstanceOwnershipResponse struct {
	Instance int `json:"instance"`
}

// ChangeInstancePermissionsResponse is the response body for PATCH /cloud/instance/permissions.
type ChangeInstancePermissionsResponse struct {
	Instance int `json:"instance"`
}

// RecoverInstanceResponse is the response body for PATCH /cloud/instance/recover.
type RecoverInstanceResponse struct {
	Instance int `json:"instance"`
}

// AddInstanceScheduleResponse is the response body for POST /cloud/instance/schedule.
type AddInstanceScheduleResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceScheduleResponse is the response body for PATCH /cloud/instance/schedule.
type UpdateInstanceScheduleResponse struct {
	Instance int `json:"instance"`
}

// AddInstanceSecurityGroupResponse is the response body for POST /cloud/instance/security-group.
type AddInstanceSecurityGroupResponse struct {
	Instance int `json:"instance"`
}

// InstancesShowbackResponse is the response body for GET /cloud/instance/showback.
type InstancesShowbackResponse struct {
	Showback []Showback `json:"showback"`
}

// ResizeInstanceResponse is the response body for PATCH /cloud/instance/size.
type ResizeInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceSnapshotResponse is the response body for POST /cloud/instance/snapshot.
type CreateInstanceSnapshotResponse struct {
	Snapshot int `json:"snapshot"`
}

// RevertInstanceSnapshotResponse is the response body for PATCH /cloud/instance/snapshot/revert.
type RevertInstanceSnapshotResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceTemplateResponse is the response body for PATCH /cloud/instance/template.
type UpdateInstanceTemplateResponse struct {
	Instance int `json:"instance"`
}

// UnlockInstanceResponse is the response body for PATCH /cloud/instance/unlock.
type UnlockInstanceResponse struct {
	Instance int `json:"instance"`
}

// InstanceResponse is the response body for GET /cloud/instance/{instance}.
type InstanceResponse struct {
	Instance LockedInstance `json:"instance"`
}

// CreateVNCProxyResponse is the response body for POST /cloud/instance/{instance}/vnc.
type CreateVNCProxyResponse struct {
	Password string `json:"password"`
	Token    string `json:"token"`
	Name     string `json:"name"`
}

//
// Security Group
//

// SecurityGroupResponse is the response body for GET /cloud/security-group.
type SecurityGroupResponse struct {
	SecurityGroup SecurityGroup `json:"security_group"`
}

// SecurityGroupsResponse is the response body for GET /cloud/security-group.
type SecurityGroupsResponse struct {
	SecurityGroups []SecurityGroup `json:"security_group"`
}

// CreateSecurityGroupResponse is the response body for POST /cloud/security-group.
type CreateSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// CloneSecurityGroupResponse is the response body for POST /cloud/security-group/clone.
type CloneSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// UpdateSecurityGroupResponse is the response body for PATCH /cloud/security-group.
type UpdateSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// CommitSecurityGroupResponse is the response body for PATCH /cloud/security-group/commit.
type CommitSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// ChangeSecurityGroupPermissionsResponse is the response body for PATCH /cloud/security-group/chmod.
type ChangeSecurityGroupPermissionsResponse struct {
	SecurityGroup int `json:"security_group"`
}

// ChangeSecurityGroupOwnershipResponse is the response body for PATCH /cloud/security-group/chown.
type ChangeSecurityGroupOwnershipResponse struct {
	SecurityGroup int `json:"security_group"`
}

// RenameSecurityGroupResponse is the response body for PATCH /cloud/security-group/rename.
type RenameSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

//
// System
//

// SystemVersionResponse is the response body for GET /cloud/system/version.
type SystemVersionResponse struct {
	Version string `json:"version"`
}

// SystemConfigResponse is the response body for GET /cloud/system/config.
type SystemConfigResponse struct {
	Config HyperCloudConfiguration `json:"config"`
}

//
// Virtual Data Center
//

// DataCenterResponse is the response body for GET /cloud/datacenter.
type DataCenterResponse struct {
	DataCenter DataCenter `json:"datacenter"`
}

// DataCentersResponse is the response body for GET /cloud/datacenter.
type DataCentersResponse struct {
	DataCenters []DataCenter `json:"datacenters"`
}

// CreateDataCenterResponse is the response body for POST /cloud/datacenter.
type CreateDataCenterResponse struct {
	DataCenter int `json:"datacenter"`
}

// UpdateDataCenterResponse is the response body for PATCH /cloud/datacenter.
type UpdateDataCenterResponse struct {
	DataCenter int `json:"datacenter"`
}

// RenameDataCenterResponse is the response body for PATCH /cloud/datacenter/name.
type RenameDataCenterResponse struct {
	DataCenter int `json:"datacenter"`
}

// AddDataCenterGroupResponse is the response body for PATCH /cloud/datacenter/group.
type AddDataCenterGroupResponse struct {
	DataCenter int `json:"datacenter"`
}

// AddDataCenterClusterResponse is the response body for PATCH /cloud/datacenter/cluster.
type AddDataCenterClusterResponse struct {
	DataCenter int `json:"datacenter"`
}

// AddDataCenterHostResponse is the response body for PATCH /cloud/datacenter/host.
type AddDataCenterHostResponse struct {
	DataCenter int `json:"datacenter"`
}

// AddDataCenterDatastoreResponse is the response body for PATCH /cloud/datacenter/datastore.
type AddDataCenterDatastoreResponse struct {
	DataCenter int `json:"datacenter"`
}

// AddDataCenterNetworkResponse is the response body for PATCH /cloud/datacenter/network.
type AddDataCenterNetworkResponse struct {
	DataCenter int `json:"datacenter"`
}

//
// Instance Group
//

// InstanceGroupResponse is the response body for GET /cloud/instance-group.
type InstanceGroupResponse struct {
	InstanceGroup InstanceGroup `json:"instance_group"`
}

// InstanceGroupsResponse is the response body for GET /cloud/instance-group.
type InstanceGroupsResponse struct {
	InstanceGroups []InstanceGroup `json:"instance_groups"`
}

// CreateInstanceGroupResponse is the response body for POST /cloud/instance-group.
type CreateInstanceGroupResponse struct {
	Group int `json:"group"`
}

// UpdateInstanceGroupResponse is the response body for PATCH /cloud/instance-group.
type UpdateInstanceGroupResponse struct {
	Group int `json:"group"`
}

// ChangeInstanceGroupOwnershipResponse is the response body for PATCH /cloud/instance-group/ownership.
type ChangeInstanceGroupOwnershipResponse struct {
	Group int `json:"group"`
}

// ChangeInstanceGroupPermissionsResponse is the response body for PATCH /cloud/instance-group/permissions.
type ChangeInstanceGroupPermissionsResponse struct {
	Group int `json:"group"`
}

// RenameInstanceGroupResponse is the response body for PATCH /cloud/instance-group/name.
type RenameInstanceGroupResponse struct {
	Group int `json:"group"`
}

// LockInstanceGroupResponse is the response body for PATCH /cloud/instance-group/lock.
type LockInstanceGroupResponse struct {
	Group int       `json:"group"`
	Time  time.Time `json:"time"`
}

// UnlockInstanceGroupResponse is the response body for PATCH /cloud/instance-group/unlock.
type UnlockInstanceGroupResponse struct {
	Group int `json:"group"`
}

//
// User
//

// UserResponse is the response body for GET /cloud/user.
type UserResponse struct {
	User User `json:"user"`
}

// UsersResponse is the response body for GET /cloud/user.
type UsersResponse struct {
	Users []User `json:"users"`
}

// UserLoginResponse is the response body for POST /cloud/user/login.
type UserLoginResponse struct {
	Token string `json:"token"`
}

// UserQuotaResponse is the response body for GET /cloud/user/quota.
type UserQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

// CreateUserResponse is the response body for POST /cloud/user.
type CreateUserResponse struct {
	User int `json:"user"`
}

// ChangeUserPasswordResponse is the response body for PATCH /cloud/user/password.
type ChangeUserPasswordResponse struct {
	User int `json:"user"`
}

// UpdateUserResponse  is the response body for PATCH /cloud/user.
type UpdateUserResponse struct {
	User int `json:"user"`
}

// ChangeUserAuthResponse is the response body for PATCH /cloud/user/auth.
type ChangeUserAuthResponse struct {
	User int `json:"user"`
}

// SetUserQuotaResponse is the response body for PATCH /cloud/user/quota.
type SetUserQuotaResponse struct {
	User int `json:"user"`
}

// ChangeUserGroupResponse is the response body for PATCH /cloud/user/group.
type ChangeUserGroupResponse struct {
	User int `json:"user"`
}

// AddUserGroupResponse is the response body for POST /cloud/user/group.
type AddUserGroupResponse struct {
	User int `json:"user"`
}

// EnableUserResponse is the response body for PATCH /cloud/user/enable.
type EnableUserResponse struct {
	User int `json:"user"`
}

// UpdateDefaultUserQuotaResponse is the response body for POST /cloud/user/quota.
type UpdateDefaultUserQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

//
// Market
//

// MarketResponse is the response body for GET /cloud/market.
type MarketResponse struct {
	Market Marketplace `json:"market"`
}

// MarketsResponse is the response body for GET /cloud/market.
type MarketsResponse struct {
	Markets []Marketplace `json:"market"`
}

// MarketAppResponse is the response body for GET /cloud/market/app.
type MarketAppResponse struct {
	App MarketplaceApp `json:"application"`
}

// MarketAppsResponse is the response body for GET /cloud/market/app.
type MarketAppsResponse struct {
	Apps []MarketplaceApp `json:"applications"`
}

// ChangeMarketAppOwnershipResponse is the response body for PATCH /cloud/market/app/ownership.
type ChangeMarketAppOwnershipResponse struct {
	MarketApp int `json:"market_app"`
}

// ChangeMarketOwnershipResponse is the response body for PATCH /cloud/market/ownership.
type ChangeMarketOwnershipResponse struct {
	Market int `json:"market"`
}

// ChangeMarketAppPermissionsResponse is the response body for PATCH /cloud/market/app/permissions.
type ChangeMarketAppPermissionsResponse struct {
	MarketApp int `json:"market_app"`
}

// ChangeMarketPermissionsResponse is the response body for PATCH /cloud/market/permissions.
type ChangeMarketPermissionsResponse struct {
	Market int `json:"market"`
}

// CreateMarketResponse is the response body for POST /cloud/market.
type CreateMarketResponse struct {
	Market int `json:"market"`
}

// CreateMarketAppResponse is the response body for POST /cloud/market/app.
type CreateMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// EnableMarketAppResponse is the response body for PATCH /cloud/market/app/enable.
type EnableMarketAppResponse struct {
	Market int `json:"market"`
}

// EnableMarketResponse is the response body for PATCH /cloud/market/enable.
type EnableMarketResponse struct {
	Market int `json:"market"`
}

// UpdateMarketAppResponse is the response body for PATCH /cloud/market/app.
type UpdateMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// UpdateMarketResponse is the response body for PATCH /cloud/market.
type UpdateMarketResponse struct {
	Market int `json:"market"`
}

// RenameMarketAppResponse is the response body for PATCH /cloud/market/app/name.
type RenameMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// RenameMarketResponse is the response body for PATCH /cloud/market/name.
type RenameMarketResponse struct {
	Market int `json:"market"`
}

// LockMarketAppResponse is the response body for PATCH /cloud/market/app/lock.
type LockMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// UnlockMarketAppResponse is the response body for PATCH /cloud/market/app/unlock.
type UnlockMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

//
// Template
//

// TemplateResponse is the response body for GET /cloud/template.
type TemplateResponse struct {
	Template InstanceTemplate `json:"template"`
}

// TemplatesResponse is the response body for GET /cloud/template.
type TemplatesResponse struct {
	Templates []InstanceTemplate `json:"template"`
}

// CreateTemplateResponse is the response body for POST /cloud/template.
type CreateTemplateResponse struct {
	Template int `json:"template"`
}

// CloneTemplateResponse is the response body for POST /cloud/template/clone.
type CloneTemplateResponse struct {
	Template int `json:"template"`
}

// DeleteTemplateResponse is the response body for DELETE /cloud/template/{template}.
type DeleteTemplateResponse struct {
	Template int `json:"template"`
}

// InstantiateTemplateResponse is the response body for DELETE /cloud/template/{template}.
type InstantiateTemplateResponse struct {
	Template int `json:"template"`
}

// UpdateTemplateResponse is the response body for PATCH /cloud/template.
type UpdateTemplateResponse struct {
	Template int `json:"template"`
}

// ChangeTemplatePermissionsResponse is the response body for PATCH /cloud/template/permissions.
type ChangeTemplatePermissionsResponse struct {
	Template int `json:"template"`
}

// ChangeTemplateOwnershipResponse is the response body for PATCH /cloud/template/ownership.
type ChangeTemplateOwnershipResponse struct {
	Template int `json:"template"`
}

// RenameTemplateResponse is the response body for PATCH /cloud/template/name.
type RenameTemplateResponse struct {
	Template int `json:"template"`
}

// LockTemplateResponse is the response body for PATCH /cloud/template/lock.
type LockTemplateResponse struct {
	Template int `json:"template"`
}

// UnlockTemplateResponse is the response body for PATCH /cloud/template/unlock.
type UnlockTemplateResponse struct {
	Template int `json:"template"`
}

//
// Network
//

// NetworkResponse is the response body for GET /cloud/network.
type NetworkResponse struct {
	Network Network `json:"network"`
}

// NetworksResponse is the response body for GET /cloud/network.
type NetworksResponse struct {
	VNets []Network `json:"networks"`
}

// NetworkTemplateResponse is the response body for GET /cloud/network.
type NetworkTemplateResponse struct {
	Template nettmpl.Template `json:"template"`
}

// NetworkTemplatesResponse is the response body for GET /cloud/network.
type NetworkTemplatesResponse struct {
	Templates []nettmpl.Template `json:"template"`
}

// CreateNetworkResponse is the response body for POST /cloud/network.
type CreateNetworkResponse struct {
	Network int `json:"network"`
}

// AddNetworkAddressRangeResponse is the response body for POST /cloud/network/address-range.
type AddNetworkAddressRangeResponse struct {
	Network int `json:"network"`
}

// UpdateNetworkAddressRangeResponse is the response body for PATCH /cloud/network/address-range.
type UpdateNetworkAddressRangeResponse struct {
	Network int `json:"network"`
}

// ReserveNetworkResponse is the response body for POST /cloud/network/reserve.
type ReserveNetworkResponse struct {
	Network int `json:"network"`
}

// HoldNetworkResponse is the response body for PATCH /cloud/network/hold.
type HoldNetworkResponse struct {
	Network int `json:"network"`
}

// ReleaseNetworkResponse is the response body for PATCH /cloud/network/release.
type ReleaseNetworkResponse struct {
	Network int `json:"network"`
}

// UpdateNetworkResponse is the response body for PATCH /cloud/network.
type UpdateNetworkResponse struct {
	Network int `json:"network"`
}

// ChangeNetworkPermissionsResponse is the response body for PATCH /cloud/network/permissions.
type ChangeNetworkPermissionsResponse struct {
	Network int `json:"network"`
}

// ChangeNetworkOwnershipResponse is the response body for PATCH /cloud/network/ownership.
type ChangeNetworkOwnershipResponse struct {
	Network int `json:"network"`
}

// RenameNetworkResponse is the response body for PATCH /cloud/network/name.
type RenameNetworkResponse struct {
	Network int `json:"network"`
}

// LockNetworkResponse is the response body for PATCH /cloud/network/lock.
type LockNetworkResponse struct {
	Network int       `json:"network"`
	Time    time.Time `json:"time"`
}

// UnlockNetworkResponse is the response body for PATCH /cloud/network/unlock.
type UnlockNetworkResponse struct {
	Network int `json:"network"`
}

// RecoverNetworkResponse is the response body for PATCH /cloud/network/recover.
type RecoverNetworkResponse struct {
	Network int `json:"network"`
}

// CreateNetworkTemplateResponse is the response body for POST /cloud/network/template.
type CreateNetworkTemplateResponse struct {
	Template int `json:"template"`
}

// CloneNetworkTemplateResponse is the response body for POST /cloud/network/template/clone.
type CloneNetworkTemplateResponse struct {
	Template int `json:"template"`
}

// InstantiateNetworkTemplateResponse is the response body for PATCH /cloud/network/template/instantiate.
type InstantiateNetworkTemplateResponse struct {
	Template int `json:"network"`
}

// UpdateNetworkTemplateResponse is the response body for PATCH /cloud/network/template.
type UpdateNetworkTemplateResponse struct {
	Template int `json:"template"`
}

// ChangeNetworkTemplateOwnershipResponse is the response body for PATCH /cloud/network/template/ownership.
type ChangeNetworkTemplateOwnershipResponse struct {
	Template int `json:"template"`
}

// ChangeNetworkTemplatePermissionsResponse is the response body for PATCH /cloud/network/template/permissions.
type ChangeNetworkTemplatePermissionsResponse struct {
	Template int `json:"template"`
}

// RenameNetworkTemplateResponse is the response body for PATCH /cloud/network/template/name.
type RenameNetworkTemplateResponse struct {
	Template int `json:"template"`
}

// LockNetworkTemplateResponse is the response body for PATCH /cloud/network/template/lock.
type LockNetworkTemplateResponse struct {
	Template int       `json:"template"`
	Time     time.Time `json:"time"`
}

// UnlockNetworkTemplateResponse is the response body for PATCH /cloud/network/template/unlock.
type UnlockNetworkTemplateResponse struct {
	Template int `json:"template"`
}

//
// VRouter
//

// RouterResponse is the response body for GET /cloud/router.
type RouterResponse struct {
	Router Router `json:"router"`
}

// RoutersResponse is the response body for GET /cloud/router.
type RoutersResponse struct {
	Routers []Router `json:"routers"`
}

// CreateRouterResponse is the response body for POST /cloud/router.
type CreateRouterResponse struct {
	Router int `json:"router"`
}

// InstantiateRouterResponse is the response body for PATCH /cloud/router/instantiate.
type InstantiateRouterResponse struct {
	Router int `json:"router"`
}

// CreateRouterNICResponse is the response body for PATCH /cloud/router/nic.
type CreateRouterNICResponse struct {
	Router int `json:"router"`
}

// UpdateRouterResponse is the response body for PATCH /cloud/router.
type UpdateRouterResponse struct {
	Router int `json:"router"`
}

// ChangeRouterPermissionsResponse is the response body for PATCH /cloud/router/permissions.
type ChangeRouterPermissionsResponse struct {
	Router int `json:"router"`
}

// ChangeRouterOwnershipResponse is the response body for PATCH /cloud/router/ownership.
type ChangeRouterOwnershipResponse struct {
	Router int `json:"router"`
}

// RenameRouterResponse is the response body for PATCH /cloud/router/name.
type RenameRouterResponse struct {
	Router int `json:"router"`
}

// LockRouterResponse is the response body for PATCH /cloud/router/lock.
type LockRouterResponse struct {
	Router int       `json:"router"`
	Time   time.Time `json:"time"`
}

// UnlockRouterResponse is the response body for PATCH /cloud/router/unlock.
type UnlockRouterResponse struct {
	Router int `json:"router"`
}

//
// Zone
//

// ZoneResponse is the response body for GET /cloud/zone.
type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

// ZonesResponse is the response body for GET /cloud/zone.
type ZonesResponse struct {
	Zones []Zone `json:"zones"`
}

// ZonesRaftStatusResponse is the response body for GET /cloud/zone/raft.
type ZonesRaftStatusResponse struct {
	Status RaftStatus `json:"status"`
}

// CreateZoneResponse is the response body for POST /cloud/zone.
type CreateZoneResponse struct {
	Zone int `json:"zone"`
}

// EnableZoneResponse is the response body for PATCH /cloud/zone/enable.
type EnableZoneResponse struct {
	Zone int `json:"zone"`
}

// UpdateZoneResponse is the response body for PATCH /cloud/zone.
type UpdateZoneResponse struct {
	Zone int `json:"zone"`
}

// RenameZoneResponse is the response body for PATCH /cloud/zone/name.
type RenameZoneResponse struct {
	Zone int `json:"zone"`
}
