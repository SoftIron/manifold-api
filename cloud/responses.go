package cloud

import (
	"time"

	"github.com/softiron/hypercloud-api/cloud/instance"
	"github.com/softiron/hypercloud-api/cloud/vnettmpl"
)

// ErrorResponse is the response body for all error responses.
type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

//
// ACL
//

// CreateACLResponse is the response body for POST /acl.
type CreateACLResponse struct {
	ACL int `json:"acl"`
}

// ListACLsResponse is the response body for GET /acl.
type ListACLsResponse struct {
	ACLs []ACL `json:"acls"`
}

//
// Cluster
//

// ListClusterResponse is the response body for GET /cluster.
type ListClusterResponse struct {
	Cluster Cluster `json:"cluster"`
}

// ListClustersResponse is the response body for GET /cluster.
type ListClustersResponse struct {
	Clusters []Cluster `json:"cluster"`
}

// CreateClusterResponse is response body for POST /cluster.
type CreateClusterResponse struct {
	Cluster int `json:"cluster"`
}

// UpdateClusterResponse is response body for PATCH /cluster.
type UpdateClusterResponse struct {
	Cluster int `json:"cluster"`
}

// AddClusterHostResponse is the response body for PATCH /cluster/host.
type AddClusterHostResponse struct {
	Host int `json:"host"`
}

// AddClusterDatastoreResponse is the response body for PATCH /cluster/datastore.
type AddClusterDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// AddClusterNetworkResponse is the response body for PATCH /cluster/vnet.
type AddClusterNetworkResponse struct {
	VNet int `json:"vnet"`
}

// RenameClusterResponse is the response body for PATCH /cluster/name.
type RenameClusterResponse struct {
	Cluster int `json:"cluster"`
}

//
// Datastore
//

// ListDatastoreResponse is the response body for GET /datastore.
type ListDatastoreResponse struct {
	Datastore Datastore `json:"datastore"`
}

// ListDatastoresResponse is the response body for GET /datastore.
type ListDatastoresResponse struct {
	Datastores []Datastore `json:"datastore"`
}

// CreateDatastoreResponse is the response body for POST /datastore.
type CreateDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// UpdateDatastoreResponse is the response body for PATCH /datastore.
type UpdateDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// ChangeDatastorePermissionsResponse is the response body for PATCH /datastore/permissions.
type ChangeDatastorePermissionsResponse struct {
	Datastore int `json:"datastore"`
}

// ChangeDatastoreOwnershipResponse is the response body for PATCH /datastore/ownership.
type ChangeDatastoreOwnershipResponse struct {
	Datastore int `json:"datastore"`
}

// RenameDatastoreResponse is the response body for PATCH /datastore/name.
type RenameDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

// EnableDatastoreResponse is the response body for PATCH /datastore/enable.
type EnableDatastoreResponse struct {
	Datastore int `json:"datastore"`
}

//
// Document
//

// ListDocumentResponse is the response body for GET /document.
type ListDocumentResponse struct {
	Document Document `json:"document"`
}

// ListDocumentsResponse is the response body for GET /document.
type ListDocumentsResponse struct {
	Documents []Document `json:"document"`
}

// AllocateDocumentResponse is the response body for POST /document.
type AllocateDocumentResponse struct {
	Document int `json:"document"`
}

// CloneDocumentResponse is the response body for POST /document/clone.
type CloneDocumentResponse struct {
	Document int `json:"document"`
}

// UpdateDocumentResponse is the response body for PATCH /document.
type UpdateDocumentResponse struct {
	Document int `json:"document"`
}

// ChangeDocumentPermissionsResponse is the response body for PATCH /document/permissions.
type ChangeDocumentPermissionsResponse struct {
	Document int `json:"document"`
}

// ChangeDocumentOwnershipResponse is the response body for PATCH /document/ownership.
type ChangeDocumentOwnershipResponse struct {
	Document int `json:"document"`
}

// RenameDocumentResponse is the response body for PATCH /document/name.
type RenameDocumentResponse struct {
	Document int `json:"document"`
}

// LockDocumentResponse is the response body for PATCH /document/lock.
type LockDocumentResponse struct {
	Document int       `json:"document"`
	Time     time.Time `json:"time"`
}

// UnlockDocumentResponse is the response body for PATCH /document/unlock.
type UnlockDocumentResponse struct {
	Document int `json:"document"`
}

//
// Group
//

// ListGroupResponse is the response body for GET /group.
type ListGroupResponse struct {
	Group Group `json:"group"`
}

// ListGroupsResponse is the response body for GET /group.
type ListGroupsResponse struct {
	Groups []Group `json:"groups"`
}

// ListGroupQuotaResponse is the response body for GET /group/quota.
type ListGroupQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

// CreateGroupResponse is the response body for POST /group.
type CreateGroupResponse struct {
	Group int `json:"group"`
}

// UpdateGroupResponse is the response body for PATCH /group.
type UpdateGroupResponse struct {
	Group int `json:"group"`
}

// UpdateGroupQuotaResponse is the response body for PATCH /group/{group}/quota.
type UpdateGroupQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quotas"`
}

// AddGroupAdminResponse is the response body for POST /group/admin.
type AddGroupAdminResponse struct {
	Group int `json:"group"`
}

// SetGroupQuotaResponse is the response body for POST /group/quota.
type SetGroupQuotaResponse struct {
	Group int `json:"group"`
}

//
// Hook
//

// ListHookResponse is the response body for GET /hook.
type ListHookResponse struct {
	Hook Hook `json:"hook"`
}

// ListHooksResponse is the response body for GET /hook.
type ListHooksResponse struct {
	Hooks []Hook `json:"hook"`
}

// ListHookLogResponse is the response body for GET /hook/log.
type ListHookLogResponse struct {
	HookLogs []HookLog `json:"hook_log"`
}

// CreateHookResponse is the response body for POST /hook.
type CreateHookResponse struct {
	Hook int `json:"hook"`
}

// UpdateHookResponse is the response body for PATCH /hook.
type UpdateHookResponse struct {
	Hook int `json:"hook"`
}

// RenameHookResponse is the response body for PATCH /hook/name.
type RenameHookResponse struct {
	Hook int `json:"hook"`
}

// LockHookResponse is the response body for PATCH /hook/lock.
type LockHookResponse struct {
	Hook int       `json:"hook"`
	Time time.Time `json:"time"`
}

// UnlockHookResponse is the response body for PATCH /hook/unlock.
type UnlockHookResponse struct {
	Hook int `json:"hook"`
}

// RetryHookResponse is the response body for PATCH /hook/retry.
type RetryHookResponse struct {
	Hook int `json:"hook"`
}

//
// Compute (v1 called this Host, but this is just the compute resources)
//

// ListComputeHostResponse is the response body for GET /compute/host.
type ListComputeHostResponse struct {
	Host ComputeHost `json:"host"`
}

// ListComputeHostsResponse is the response body for GET /compute/host.
type ListComputeHostsResponse struct {
	Hosts []ComputeHost `json:"hosts"`
}

// ListComputeHostMonitoringResponse is the response body for GET /compute/host/monitoring.
type ListComputeHostMonitoringResponse struct {
	Monitoring []HostMonitoring `json:"monitoring"`
}

// ListComputeHostsMonitoringResponse is the response body for GET /compute/host/monitoring.
type ListComputeHostsMonitoringResponse struct {
	Monitoring []HostMonitoring `json:"monitoring"`
}

// CreateComputeHostResponse is the response body for POST /compute/host.
type CreateComputeHostResponse struct {
	Host int `json:"host"`
}

// SetComputeHostStatusResponse is the response body for PATCH /compute/host/status.
type SetComputeHostStatusResponse struct {
	Host int `json:"host"`
}

// UpdateComputeHostResponse is the response body for PATCH /compute/host.
type UpdateComputeHostResponse struct {
	Host int `json:"host"`
}

// RenameComputeHostResponse is the response body for PATCH /compute/host/name.
type RenameComputeHostResponse struct {
	Host int `json:"host"`
}

//
// Host
//

// ListHostsResponse is the response body for GET /host.
type ListHostsResponse struct {
	Hosts []Host `json:"host"`
}

//
// Image
//

// ListImageResponse is the response body for GET /image.
type ListImageResponse struct {
	Image Image `json:"image"`
}

// ListImagesResponse is the response body for GET /image.
type ListImagesResponse struct {
	Images []Image `json:"images"`
}

// CreateImageResponse is the response body for POST /image.
type CreateImageResponse struct {
	Image int `json:"image"`
}

// CloneImageResponse is the response body for POST /image/clone.
type CloneImageResponse struct {
	Image int `json:"image"`
}

// EnableImageResponse is the response body for PATCH /image/enable.
type EnableImageResponse struct {
	Image int `json:"image"`
}

// SetImagePersistentResponse is the response body for PATCH /image/persistent.
type SetImagePersistentResponse struct {
	Image int `json:"image"`
}

// ChangeImageTypeResponse is the response body for PATCH /image/type.
type ChangeImageTypeResponse struct {
	Image int `json:"image"`
}

// UpdateImageResponse is the response body for PATCH /image.
type UpdateImageResponse struct {
	Image int `json:"image"`
}

// ChangeImagePermissionsResponse is the response body for PATCH /image/permissions.
type ChangeImagePermissionsResponse struct {
	Image int `json:"image"`
}

// ChangeImageOwnershipResponse is the response body for PATCH /image/ownership.
type ChangeImageOwnershipResponse struct {
	Image int `json:"image"`
}

// RenameImageResponse is the response body for PATCH /image/name.
type RenameImageResponse struct {
	Image int `json:"image"`
}

// RevertImageSnapshotResponse is the response body for PATCH /image/snapshot/revert.
type RevertImageSnapshotResponse struct {
	Image int `json:"image"`
}

// FlattenImageSnapshotResponse is the response body for PATCH /image/snapshot/flatten.
type FlattenImageSnapshotResponse struct {
	Image int `json:"image"`
}

// LockImageResponse is the response body for PATCH /image/lock.
type LockImageResponse struct {
	Image int       `json:"image"`
	Time  time.Time `json:"time"`
}

// UnlockImageResponse is the response body for PATCH /image/unlock.
type UnlockImageResponse struct {
	Image int `json:"image"`
}

//
// Instance
//

// ListInstancesResponse is the response body for GET /instance.
type ListInstancesResponse struct {
	Instances []LockedInstance `json:"instances"`
}

// CreateInstanceResponse is the response body for POST /instance.
type CreateInstanceResponse struct {
	Instance int `json:"instance"`
}

// ListInstancesAccountingResponse is the response body for GET /instance/accounting.
type ListInstancesAccountingResponse struct {
	Accounting []AcctHistory `json:"accounting"`
}

// SetInstanceActionResponse is the response body for POST /instance/action.
type SetInstanceActionResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceConfigResponse is the response body for PATCH /instance/config.
type UpdateInstanceConfigResponse struct {
	Instance int `json:"instance"`
}

// DeployInstanceResponse is the response body for PATCH /instance/deploy.
type DeployInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskResponse is the response body for POST /instance/disk.
type CreateInstanceDiskResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskImageResponse is the response body for POST /instance/disk/image.
type CreateInstanceDiskImageResponse struct {
	Image int `json:"image"`
}

// ResizeInstanceDiskResponse is the response body for PATCH /instance/disk/size.
type ResizeInstanceDiskResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceDiskSnapshotResponse is the response body for POST /instance/disk/snapshot.
type CreateInstanceDiskSnapshotResponse struct {
	Instance int `json:"instance"`
}

// RenameInstanceDiskSnapshotResponse is the response body for PATCH /instance/disk/snapshot/name.
type RenameInstanceDiskSnapshotResponse struct {
	Instance int `json:"instance"`
}

// RevertInstanceDiskSnapshotResponse is the response body for PATCH /instance/disk/snapshot/revert.
type RevertInstanceDiskSnapshotResponse struct {
	Snapshot int `json:"snapshot"`
}

// LockInstanceResponse is the response body for PATCH /instance/lock.
type LockInstanceResponse struct {
	Instance int       `json:"instance"`
	Time     time.Time `json:"time"`
}

// ListInstancesMonitoringResponse is the response body for GET /instance/monitoring.
type ListInstancesMonitoringResponse struct {
	Monitoring []instance.Monitoring `json:"monitoring"`
}

// ListInstanceMonitoringResponse is the response body for GET /instance/monitoring/{id}.
type ListInstanceMonitoringResponse struct {
	Monitoring []instance.Monitoring `json:"monitoring"`
}

// MoveInstanceResponse is the response body for PATCH /instance/move.
type MoveInstanceResponse struct {
	Instance int `json:"instance"`
}

// RenameInstanceResponse is the response body for PATCH /instance/name.
type RenameInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceNICResponse is the response body for POST /instance/nic.
type CreateInstanceNICResponse struct {
	Instance int `json:"instance"` // TODO: would make more sense as NIC ID, maybe docs are wrong
}

// ChangeInstanceOwnershipResponse is the response body for PATCH /instance/ownership.
type ChangeInstanceOwnershipResponse struct {
	Instance int `json:"instance"`
}

// ChangeInstancePermissionsResponse is the response body for PATCH /instance/permissions.
type ChangeInstancePermissionsResponse struct {
	Instance int `json:"instance"`
}

// RecoverInstanceResponse is the response body for PATCH /instance/recover.
type RecoverInstanceResponse struct {
	Instance int `json:"instance"`
}

// AddInstanceScheduleResponse is the response body for POST /instance/schedule.
type AddInstanceScheduleResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceScheduleResponse is the response body for PATCH /instance/schedule.
type UpdateInstanceScheduleResponse struct {
	Instance int `json:"instance"`
}

// AddInstanceSecurityGroupResponse is the response body for POST /instance/security-group.
type AddInstanceSecurityGroupResponse struct {
	Instance int `json:"instance"`
}

// ListInstancesShowbackResponse is the response body for GET /instance/showback.
type ListInstancesShowbackResponse struct {
	Showback []Showback `json:"showback"`
}

// ResizeInstanceResponse is the response body for PATCH /instance/size.
type ResizeInstanceResponse struct {
	Instance int `json:"instance"`
}

// CreateInstanceSnapshotResponse is the response body for POST /instance/snapshot.
type CreateInstanceSnapshotResponse struct {
	Snapshot int `json:"snapshot"`
}

// RevertInstanceSnapshotResponse is the response body for PATCH /instance/snapshot/revert.
type RevertInstanceSnapshotResponse struct {
	Instance int `json:"instance"`
}

// UpdateInstanceTemplateResponse is the response body for PATCH /instance/template.
type UpdateInstanceTemplateResponse struct {
	Instance int `json:"instance"`
}

// UnlockInstanceResponse is the response body for PATCH /instance/unlock.
type UnlockInstanceResponse struct {
	Instance int `json:"instance"`
}

// ListInstanceResponse is the response body for GET /instance/{instance}.
type ListInstanceResponse struct {
	Instance LockedInstance `json:"instance"`
}

// PingResponse is the response from POST /ping.
type PingResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// CreateVNCProxyResponse is the response body for POST /instance/{instance}/vnc.
type CreateVNCProxyResponse struct {
	Password string `json:"password"`
	Token    string `json:"token"`
	Name     string `json:"name"`
}

//
// Security Group
//

// ListSecurityGroupResponse is the response body for GET /sg.
type ListSecurityGroupResponse struct {
	SecurityGroup SecurityGroup `json:"security_group"`
}

// ListSecurityGroupsResponse is the response body for GET /sg.
type ListSecurityGroupsResponse struct {
	SecurityGroups []SecurityGroup `json:"security_group"`
}

// CreateSecurityGroupResponse is the response body for POST /sg.
type CreateSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// CloneSecurityGroupResponse is the response body for POST /sg/clone.
type CloneSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// UpdateSecurityGroupResponse is the response body for PATCH /sg.
type UpdateSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// CommitSecurityGroupResponse is the response body for PATCH /sg/commit.
type CommitSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

// ChangeSecurityGroupPermissionsResponse is the response body for PATCH /sg/chmod.
type ChangeSecurityGroupPermissionsResponse struct {
	SecurityGroup int `json:"security_group"`
}

// ChangeSecurityGroupOwnershipResponse is the response body for PATCH /sg/chown.
type ChangeSecurityGroupOwnershipResponse struct {
	SecurityGroup int `json:"security_group"`
}

// RenameSecurityGroupResponse is the response body for PATCH /sg/rename.
type RenameSecurityGroupResponse struct {
	SecurityGroup int `json:"security_group"`
}

//
// System
//

// ListSystemVersionResponse is the response body for GET /system/version.
type ListSystemVersionResponse struct {
	Version string `json:"version"`
}

// ListSystemConfigResponse is the response body for GET /system/config.
type ListSystemConfigResponse struct {
	Config HyperCloudConfiguration `json:"config"`
}

//
// VDC
//

// ListVDCResponse is the response body for GET /vdc.
type ListVDCResponse struct {
	VDC VDC `json:"vdc"`
}

// ListVDCsResponse is the response body for GET /vdc.
type ListVDCsResponse struct {
	VDCs []VDC `json:"vdcs"`
}

// CreateVDCResponse is the response body for POST /vdc.
type CreateVDCResponse struct {
	VDC int `json:"vdc"`
}

// UpdateVDCResponse is the response body for PATCH /vdc.
type UpdateVDCResponse struct {
	VDC int `json:"vdc"`
}

// RenameVDCResponse is the response body for PATCH /vdc/name.
type RenameVDCResponse struct {
	VDC int `json:"vdc"`
}

// AddVDCGroupResponse is the response body for PATCH /vdc/group.
type AddVDCGroupResponse struct {
	VDC int `json:"vdc"`
}

// AddVDCClusterResponse is the response body for PATCH /vdc/cluster.
type AddVDCClusterResponse struct {
	VDC int `json:"vdc"`
}

// AddVDCHostResponse is the response body for PATCH /vdc/host.
type AddVDCHostResponse struct {
	VDC int `json:"vdc"`
}

// AddVDCDatastoreResponse is the response body for PATCH /vdc/datastore.
type AddVDCDatastoreResponse struct {
	VDC int `json:"vdc"`
}

// AddVDCNetworkResponse is the response body for PATCH /vdc/network.
type AddVDCNetworkResponse struct {
	VDC int `json:"vdc"`
}

//
// Instance Group
//

// ListInstanceGroupResponse is the response body for GET /instance-group.
type ListInstanceGroupResponse struct {
	InstanceGroup InstanceGroup `json:"instance_group"`
}

// ListInstanceGroupsResponse is the response body for GET /instance-group.
type ListInstanceGroupsResponse struct {
	InstanceGroups []InstanceGroup `json:"instance_groups"`
}

// CreateInstanceGroupResponse is the response body for POST /instance-group.
type CreateInstanceGroupResponse struct {
	Group int `json:"group"`
}

// UpdateInstanceGroupResponse is the response body for PATCH /instance-group.
type UpdateInstanceGroupResponse struct {
	Group int `json:"group"`
}

// ChangeInstanceGroupOwnershipResponse is the response body for PATCH /instance-group/ownership.
type ChangeInstanceGroupOwnershipResponse struct {
	Group int `json:"group"`
}

// ChangeInstanceGroupPermissionsResponse is the response body for PATCH /instance-group/permissions.
type ChangeInstanceGroupPermissionsResponse struct {
	Group int `json:"group"`
}

// RenameInstanceGroupResponse is the response body for PATCH /instance-group/name.
type RenameInstanceGroupResponse struct {
	Group int `json:"group"`
}

// LockInstanceGroupResponse is the response body for PATCH /instance-group/lock.
type LockInstanceGroupResponse struct {
	Group int       `json:"group"`
	Time  time.Time `json:"time"`
}

// UnlockInstanceGroupResponse is the response body for PATCH /instance-group/unlock.
type UnlockInstanceGroupResponse struct {
	Group int `json:"group"`
}

//
// User
//

// ListUserResponse is the response body for GET /user.
type ListUserResponse struct {
	User User `json:"user"`
}

// ListUsersResponse is the response body for GET /user.
type ListUsersResponse struct {
	Users []User `json:"users"`
}

// UserLoginResponse is the response body for POST /user/login.
type UserLoginResponse struct {
	Token string `json:"token"`
}

// ListUserQuotaResponse is the response body for GET /user/quota.
type ListUserQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

// CreateUserResponse is the response body for POST /user.
type CreateUserResponse struct {
	User int `json:"user"`
}

// ChangeUserPasswordResponse is the response body for PATCH /user/password.
type ChangeUserPasswordResponse struct {
	User int `json:"user"`
}

// UpdateUserResponse  is the response body for PATCH /user.
type UpdateUserResponse struct {
	User int `json:"user"`
}

// ChangeUserAuthResponse is the response body for PATCH /user/auth.
type ChangeUserAuthResponse struct {
	User int `json:"user"`
}

// SetUserQuotaResponse is the response body for PATCH /user/quota.
type SetUserQuotaResponse struct {
	User int `json:"user"`
}

// ChangeUserGroupResponse is the response body for PATCH /user/group.
type ChangeUserGroupResponse struct {
	User int `json:"user"`
}

// AddUserGroupResponse is the response body for POST /user/group.
type AddUserGroupResponse struct {
	User int `json:"user"`
}

// EnableUserResponse is the response body for PATCH /user/enable.
type EnableUserResponse struct {
	User int `json:"user"`
}

// UpdateDefaultUserQuotaResponse is the response body for POST /user/quota.
type UpdateDefaultUserQuotaResponse struct {
	Quota UserDefaultQuotas `json:"quota"`
}

//
// Market
//

// ListMarketResponse is the response body for GET /market.
type ListMarketResponse struct {
	Market Marketplace `json:"market"`
}

// ListMarketsResponse is the response body for GET /market.
type ListMarketsResponse struct {
	Markets []Marketplace `json:"market"`
}

// ListMarketAppResponse is the response body for GET /market/app.
type ListMarketAppResponse struct {
	App MarketplaceApp `json:"application"`
}

// ListMarketAppsResponse is the response body for GET /market/app.
type ListMarketAppsResponse struct {
	Apps []MarketplaceApp `json:"applications"`
}

// ChangeMarketAppOwnershipResponse is the response body for PATCH /market/app/ownership.
type ChangeMarketAppOwnershipResponse struct {
	MarketApp int `json:"market_app"`
}

// ChangeMarketOwnershipResponse is the response body for PATCH /market/ownership.
type ChangeMarketOwnershipResponse struct {
	Market int `json:"market"`
}

// ChangeMarketAppPermissionsResponse is the response body for PATCH /market/app/permissions.
type ChangeMarketAppPermissionsResponse struct {
	MarketApp int `json:"market_app"`
}

// ChangeMarketPermissionsResponse is the response body for PATCH /market/permissions.
type ChangeMarketPermissionsResponse struct {
	Market int `json:"market"`
}

// CreateMarketResponse is the response body for POST /market.
type CreateMarketResponse struct {
	Market int `json:"market"`
}

// CreateMarketAppResponse is the response body for POST /market/app.
type CreateMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// EnableMarketAppResponse is the response body for PATCH /market/app/enable.
type EnableMarketAppResponse struct {
	Market int `json:"market"`
}

// EnableMarketResponse is the response body for PATCH /market/enable.
type EnableMarketResponse struct {
	Market int `json:"market"`
}

// UpdateMarketAppResponse is the response body for PATCH /market/app.
type UpdateMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// UpdateMarketResponse is the response body for PATCH /market.
type UpdateMarketResponse struct {
	Market int `json:"market"`
}

// RenameMarketAppResponse is the response body for PATCH /market/app/name.
type RenameMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// RenameMarketResponse is the response body for PATCH /market/name.
type RenameMarketResponse struct {
	Market int `json:"market"`
}

// LockMarketAppResponse is the response body for PATCH /market/app/lock.
type LockMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

// UnlockMarketAppResponse is the response body for PATCH /market/app/unlock.
type UnlockMarketAppResponse struct {
	MarketApp int `json:"market_app"`
}

//
// Template
//

// ListTemplateResponse is the response body for GET /template.
type ListTemplateResponse struct {
	Template InstanceTemplate `json:"template"`
}

// ListTemplatesResponse is the response body for GET /template.
type ListTemplatesResponse struct {
	Templates []InstanceTemplate `json:"template"`
}

// CreateTemplateResponse is the response body for POST /template.
type CreateTemplateResponse struct {
	Template int `json:"template"`
}

// CloneTemplateResponse is the response body for POST /template/clone.
type CloneTemplateResponse struct {
	Template int `json:"template"`
}

// DeleteTemplateResponse is the response body for DELETE /template/{template}.
type DeleteTemplateResponse struct {
	Template int `json:"template"`
}

// InstantiateTemplateResponse is the response body for DELETE /template/{template}.
type InstantiateTemplateResponse struct {
	Template int `json:"template"`
}

// UpdateTemplateResponse is the response body for PATCH /template.
type UpdateTemplateResponse struct {
	Template int `json:"template"`
}

// ChangeTemplatePermissionsResponse is the response body for PATCH /template/permissions.
type ChangeTemplatePermissionsResponse struct {
	Template int `json:"template"`
}

// ChangeTemplateOwnershipResponse is the response body for PATCH /template/ownership.
type ChangeTemplateOwnershipResponse struct {
	Template int `json:"template"`
}

// RenameTemplateResponse is the response body for PATCH /template/name.
type RenameTemplateResponse struct {
	Template int `json:"template"`
}

// LockTemplateResponse is the response body for PATCH /template/lock.
type LockTemplateResponse struct {
	Template int `json:"template"`
}

// UnlockTemplateResponse is the response body for PATCH /template/unlock.
type UnlockTemplateResponse struct {
	Template int `json:"template"`
}

//
// VNet
//

// ListVNetResponse is the response body for GET /vnet.
type ListVNetResponse struct {
	VNet VNet `json:"vnet"`
}

// ListVNetsResponse is the response body for GET /vnet.
type ListVNetsResponse struct {
	VNets []VNet `json:"vnets"`
}

// ListVNetTemplateResponse is the response body for GET /vnet.
type ListVNetTemplateResponse struct {
	Template vnettmpl.Template `json:"template"`
}

// ListVNetTemplatesResponse is the response body for GET /vnet.
type ListVNetTemplatesResponse struct {
	Templates []vnettmpl.Template `json:"template"`
}

// CreateVNetResponse is the response body for POST /vnet.
type CreateVNetResponse struct {
	VNet int `json:"vnet"`
}

// AddVNetAddressRangeResponse is the response body for POST /vnet/address-range.
type AddVNetAddressRangeResponse struct {
	VNet int `json:"vnet"`
}

// UpdateVNetAddressRangeResponse is the response body for PATCH /vnet/address-range.
type UpdateVNetAddressRangeResponse struct {
	VNet int `json:"vnet"`
}

// ReserveVNetResponse is the response body for POST /vnet/reserve.
type ReserveVNetResponse struct {
	VNet int `json:"vnet"`
}

// HoldVNetResponse is the response body for PATCH /vnet/hold.
type HoldVNetResponse struct {
	VNet int `json:"vnet"`
}

// ReleaseVNetResponse is the response body for PATCH /vnet/release.
type ReleaseVNetResponse struct {
	VNet int `json:"vnet"`
}

// UpdateVNetResponse is the response body for PATCH /vnet.
type UpdateVNetResponse struct {
	VNet int `json:"vnet"`
}

// ChangeVNetPermissionsResponse is the response body for PATCH /vnet/permissions.
type ChangeVNetPermissionsResponse struct {
	VNet int `json:"vnet"`
}

// ChangeVNetOwnershipResponse is the response body for PATCH /vnet/ownership.
type ChangeVNetOwnershipResponse struct {
	VNet int `json:"vnet"`
}

// RenameVNetResponse is the response body for PATCH /vnet/name.
type RenameVNetResponse struct {
	VNet int `json:"vnet"`
}

// LockVNetResponse is the response body for PATCH /vnet/lock.
type LockVNetResponse struct {
	VNet int       `json:"vnet"`
	Time time.Time `json:"time"`
}

// UnlockVNetResponse is the response body for PATCH /vnet/unlock.
type UnlockVNetResponse struct {
	VNet int `json:"vnet"`
}

// RecoverVNetResponse is the response body for PATCH /vnet/recover.
type RecoverVNetResponse struct {
	VNet int `json:"vnet"`
}

// CreateVNetTemplateResponse is the response body for POST /vnet/template.
type CreateVNetTemplateResponse struct {
	Template int `json:"template"`
}

// CloneVNetTemplateResponse is the response body for POST /vnet/template/clone.
type CloneVNetTemplateResponse struct {
	Template int `json:"template"`
}

// InstantiateVNetTemplateResponse is the response body for PATCH /vnet/template/instantiate.
type InstantiateVNetTemplateResponse struct {
	Template int `json:"vnet"`
}

// UpdateVNetTemplateResponse is the response body for PATCH /vnet/template.
type UpdateVNetTemplateResponse struct {
	Template int `json:"template"`
}

// ChangeVNetTemplateOwnershipResponse is the response body for PATCH /vnet/template/ownership.
type ChangeVNetTemplateOwnershipResponse struct {
	Template int `json:"template"`
}

// ChangeVNetTemplatePermissionsResponse is the response body for PATCH /vnet/template/permissions.
type ChangeVNetTemplatePermissionsResponse struct {
	Template int `json:"template"`
}

// RenameVNetTemplateResponse is the response body for PATCH /vnet/template/name.
type RenameVNetTemplateResponse struct {
	Template int `json:"template"`
}

// LockVNetTemplateResponse is the response body for PATCH /vnet/template/lock.
type LockVNetTemplateResponse struct {
	Template int       `json:"template"`
	Time     time.Time `json:"time"`
}

// UnlockVNetTemplateResponse is the response body for PATCH /vnet/template/unlock.
type UnlockVNetTemplateResponse struct {
	Template int `json:"template"`
}

//
// VRouter
//

// ListVRouterResponse is the response body for GET /vrouter.
type ListVRouterResponse struct {
	VRouter VRouter `json:"vrouter"`
}

// ListVRoutersResponse is the response body for GET /vrouter.
type ListVRoutersResponse struct {
	VRouters []VRouter `json:"vrouters"`
}

// CreateVRouterResponse is the response body for POST /vrouter.
type CreateVRouterResponse struct {
	Router int `json:"router"`
}

// InstantiateVRouterResponse is the response body for PATCH /vrouter/instantiate.
type InstantiateVRouterResponse struct {
	Router int `json:"router"`
}

// CreateVRouterNICResponse is the response body for PATCH /vrouter/nic.
type CreateVRouterNICResponse struct {
	Router int `json:"router"`
}

// UpdateVRouterResponse is the response body for PATCH /vrouter.
type UpdateVRouterResponse struct {
	Router int `json:"router"`
}

// ChangeVRouterPermissionsResponse is the response body for PATCH /vrouter/permissions.
type ChangeVRouterPermissionsResponse struct {
	Router int `json:"router"`
}

// ChangeVRouterOwnershipResponse is the response body for PATCH /vrouter/ownership.
type ChangeVRouterOwnershipResponse struct {
	Router int `json:"router"`
}

// RenameVRouterResponse is the response body for PATCH /vrouter/name.
type RenameVRouterResponse struct {
	Router int `json:"router"`
}

// LockVRouterResponse is the response body for PATCH /vrouter/lock.
type LockVRouterResponse struct {
	Router int       `json:"router"`
	Time   time.Time `json:"time"`
}

// UnlockVRouterResponse is the response body for PATCH /vrouter/unlock.
type UnlockVRouterResponse struct {
	Router int `json:"router"`
}

//
// Zone
//

// ListZoneResponse is the response body for GET /zone.
type ListZoneResponse struct {
	Zone Zone `json:"zone"`
}

// ListZonesResponse is the response body for GET /zone.
type ListZonesResponse struct {
	Zones []Zone `json:"zones"`
}

// ListZonesRaftStatusResponse is the response body for GET /zone/raft.
type ListZonesRaftStatusResponse struct {
	Status RaftStatus `json:"status"`
}

// CreateZoneResponse is the response body for POST /zone.
type CreateZoneResponse struct {
	Zone int `json:"zone"`
}

// EnableZoneResponse is the response body for PATCH /zone/enable.
type EnableZoneResponse struct {
	Zone int `json:"zone"`
}

// UpdateZoneResponse is the response body for PATCH /zone.
type UpdateZoneResponse struct {
	Zone int `json:"zone"`
}

// RenameZoneResponse is the response body for PATCH /zone/name.
type RenameZoneResponse struct {
	Zone int `json:"zone"`
}
