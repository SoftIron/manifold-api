package cloud

import "github.com/softiron/manifold-api/deprecated/v1/hc"

//
// ACL
//

// CreateACLRequest is the request body for POST /cloud/acl.
type CreateACLRequest struct {
	User     int  `json:"user"`
	Resource int  `json:"resource"`
	Rights   int  `json:"rights"`
	Zone     *int `json:"zone"`
}

//
// Cluster
//

// CreateClusterRequest is request body for POST /cloud/cluster.
type CreateClusterRequest struct {
	Name string `json:"name"`
}

// UpdateClusterRequest is request body for PATCH /cloud/cluster/{cluster}.
type UpdateClusterRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameClusterRequest is the request body for PATCH /cloud/cluster/{cluster}/name.
type RenameClusterRequest struct {
	Name string `json:"name"`
}

//
// Datastore
//

// CreateDatastoreRequest is the request body for POST /cloud/datastore.
type CreateDatastoreRequest struct {
	Template string `json:"template"`
	Cluster  *int   `json:"cluster"`
}

// UpdateDatastoreRequest is the request body for PATCH /cloud/datastore/{datastore}.
type UpdateDatastoreRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeDatastorePermissionsRequest is the request body for PATCH /cloud/datastore/{datastore}/permissions.
type ChangeDatastorePermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeDatastoreOwnershipRequest is the request body for PATCH /cloud/datastore/{datastore}/ownership.
type ChangeDatastoreOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameDatastoreRequest is the request body for PATCH /cloud/datastore/{datastore}/name.
type RenameDatastoreRequest struct {
	Name string `json:"name"`
}

// EnableDatastoreRequest is the request body for PATCH /cloud/datastore/{datastore}/enable.
type EnableDatastoreRequest struct {
	Enable bool `json:"enable"`
}

//
// Document
//

// AllocateDocumentRequest is the request body for POST /cloud/document.
type AllocateDocumentRequest struct {
	Template string `json:"template"`
	Type     int    `json:"type"`
}

// CloneDocumentRequest is the request body for POST /cloud/document/{document}/clone.
type CloneDocumentRequest struct {
	Name string `json:"name"`
}

// UpdateDocumentRequest is the request body for PATCH /cloud/document/{document}.
type UpdateDocumentRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeDocumentPermissionsRequest is the request body for PATCH /cloud/document/{document}/permissions.
type ChangeDocumentPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeDocumentOwnershipRequest is the request body for PATCH /cloud/document/{document}/ownership.
type ChangeDocumentOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameDocumentRequest is the request body for PATCH /cloud/document/{document}/name.
type RenameDocumentRequest struct {
	Name string `json:"name"`
}

// LockDocumentRequest is the request body for PATCH /cloud/document/{document}/lock.
type LockDocumentRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Group
//

// CreateGroupRequest is the request body for POST /cloud/group.
type CreateGroupRequest struct {
	Name string `json:"name"`
}

// UpdateGroupRequest is the request body for PATCH /cloud/group.
type UpdateGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *UpdateGroupRequest) SetTemplate(t GroupTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// AddGroupAdminRequest is the request body for POST /cloud/group/admin.
type AddGroupAdminRequest struct {
	User int `json:"user"`
}

// SetGroupQuotaRequest is the request body for POST /cloud/group/quota.
type SetGroupQuotaRequest struct {
	Template string `json:"template"`
}

// UpdateGroupQuotaRequest is the request body for PATCH /cloud/group/{group}/quota.
type UpdateGroupQuotaRequest struct {
	Template string `json:"template"`
}

//
// Hook
//

// CreateHookRequest is the request body for POST /cloud/hook.
type CreateHookRequest struct {
	Template string `json:"template"`
}

// UpdateHookRequest is the request body for PATCH /cloud/hook/{hook}.
type UpdateHookRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameHookRequest is the request body for PATCH /cloud/hook/{hook}/name.
type RenameHookRequest struct {
	Name string `json:"name"`
}

// LockHookRequest is the request body for PATCH /cloud/hook/{hook}/lock.
type LockHookRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

// UnlockHookRequest is the request body for PATCH /cloud/hook/{hook}/unlock.
type UnlockHookRequest struct {
	Lock int `json:"lock"`
}

// RetryHookRequest is the request body for PATCH /cloud/hook/{hook}/retry.
type RetryHookRequest struct {
	Execution int `json:"execution"`
}

//
// Compute (v1 called this Host, but this is just the compute resources)
//

// CreateHostRequest is the request body for POST /cloud/host.
type CreateHostRequest struct {
	Hostname    string `json:"hostname"`
	InfoManager string `json:"info_manager"`
	VMManager   string `json:"vm_manager"`
	Cluster     *int   `json:"cluster"`
}

// SetHostStatusRequest is the request body for PATCH /cloud/host/{host}/status.
type SetHostStatusRequest struct {
	Status string `json:"status"`
}

// UpdateHostRequest is the request body for PATCH /cloud/host/{host}.
type UpdateHostRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameHostRequest is the request body for PATCH /cloud/host/{host}/name.
type RenameHostRequest struct {
	Hostname string `json:"hostname"`
}

//
// Image
//

// CreateImageRequest is the request body for POST /cloud/image.
type CreateImageRequest struct {
	Template        string `json:"template"`
	Datastore       int    `json:"datastore"`
	EnforceCapacity bool   `json:"enforce_capacity"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *CreateImageRequest) SetTemplate(t ImageTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// CloneImageRequest is the request body for POST /cloud/image/clone.
type CloneImageRequest struct {
	Name      string `json:"name"`
	Datastore *int   `json:"datastore"`
}

// EnableImageRequest is the request body for PATCH /cloud/image/{image}/enable.
type EnableImageRequest struct {
	Enable bool `json:"enable"`
}

// SetImagePersistentRequest is the request body for PATCH /cloud/image/{image}/persistent.
type SetImagePersistentRequest struct {
	Persistent bool `json:"persistent"`
}

// ChangeImageTypeRequest is the request body for PATCH /cloud/image/{image}/type.
type ChangeImageTypeRequest struct {
	Type string `json:"type"`
}

// UpdateImageRequest is the request body for PATCH /cloud/image/{image}.
type UpdateImageRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *UpdateImageRequest) SetTemplate(t ImageTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// ChangeImagePermissionsRequest is the request body for PATCH /cloud/image/{image}/permissions.
type ChangeImagePermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeImageOwnershipRequest is the request body for PATCH /cloud/image/{image}/ownership.
type ChangeImageOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameImageRequest is the request body for PATCH /cloud/image/{image}/name.
type RenameImageRequest struct {
	Name string `json:"name"`
}

// LockImageRequest is the request body for PATCH /cloud/image/{image}/lock.
type LockImageRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Instance
//

// CreateInstanceRequest is the request body for POST /cloud/instance.
type CreateInstanceRequest struct {
	Template string `json:"template"`
	Pending  bool   `json:"pending"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *CreateInstanceRequest) SetTemplate(t InstanceTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// SetInstanceActionRequest is the request body for PATCH /cloud/instance/{instance}/action.
type SetInstanceActionRequest struct {
	// Action to perform
	Action string `json:"action" enums:"hold,poweroff,poweroff-hard,reboot,reboot-hard,release,resched,resume,stop,suspend,terminate,terminate-hard,undeploy,undeploy-hard"`
}

// UpdateInstanceConfigRequest is the request body for PATCH /cloud/instance/{instance}/config.
type UpdateInstanceConfigRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// DeployInstanceRequest is the request body for POST /cloud/instance/{instance}/deploy.
type DeployInstanceRequest struct {
	Host      int    `json:"host"`
	Enforce   bool   `json:"enforce"`
	Datastore *int   `json:"datastore"`
	Template  string `json:"template"`
}

// CreateInstanceDiskRequest is the request body for POST /cloud/instance/disk.
type CreateInstanceDiskRequest struct {
	DiskTemplate string `json:"disk"`
}

// CreateInstanceDiskImageRequest is the request body for POST /cloud/instance/disk/image.
type CreateInstanceDiskImageRequest struct {
	Name      string `json:"name"`
	ImageType string `json:"image_type"`
	Snapshot  *int   `json:"snapshot"`
}

// ResizeInstanceDiskRequest is the request body for POST /cloud/instance/{instance}/disk/{disk}/size.
type ResizeInstanceDiskRequest struct {
	Size int `json:"size"`
}

// CreateInstanceDiskSnapshotRequest is the request body for POST /cloud/instance/disk/snapshot.
type CreateInstanceDiskSnapshotRequest struct {
	Description string `json:"description"`
}

// RenameInstanceDiskSnapshotRequest is the request body for PATCH /cloud/instance/{instance}/disk/{disk}/snapshot/{snapshot}/name.
type RenameInstanceDiskSnapshotRequest struct {
	Name string `json:"name"`
}

// LockInstanceRequest is the request body for PATCH /cloud/instance/{instance}/lock.
type LockInstanceRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

// MoveInstanceRequest is the request body for PATCH /cloud/instance/{instance}/move.
type MoveInstanceRequest struct {
	Host            int    `json:"host"`
	LiveMigrate     bool   `json:"live_migrate"`
	EnforceCapacity bool   `json:"enforce_capacity"` // TODO: probably should reverse this logic
	MigrationType   string `json:"migration_type"`
}

// RenameInstanceRequest is the request body for PATCH /cloud/instance/{instance}/name.
type RenameInstanceRequest struct {
	Name string `json:"name"`
}

// CreateInstanceNICRequest is the request body for POST /cloud/instance/nic.
type CreateInstanceNICRequest struct {
	NICTemplate string `json:"nic"`
}

// ChangeInstanceOwnershipRequest is the request body for PATCH /cloud/instance/{instance}/ownership.
type ChangeInstanceOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// ChangeInstancePermissionsRequest is the request body for PATCH /cloud/instance/{instance}/permissions.
type ChangeInstancePermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// RecoverInstanceRequest is the request body for PATCH /cloud/instance/{instance}/recover.
type RecoverInstanceRequest struct {
	Operation string `json:"operation"`
}

// AddInstanceScheduleRequest is the request body for POST /cloud/instance/schedule.
type AddInstanceScheduleRequest struct {
	Template string `json:"template"`
}

// UpdateInstanceScheduleRequest is the request body for PATCH /cloud/instance/{instance}/schedule.
type UpdateInstanceScheduleRequest struct {
	Action   int    `json:"action"`
	Template string `json:"template"`
}

// CalculateInstancesShowbackRequest is the request body for POST /cloud/instance/showback.
type CalculateInstancesShowbackRequest struct {
	Month Period `json:"month"`
	Year  Period `json:"year"`
}

// ResizeInstanceRequest is the request body for PATCH /cloud/instance/{instance}/size.
type ResizeInstanceRequest struct {
	Template        string `json:"template"`
	EnforceCapacity bool   `json:"enforce_capacity"` // TODO: probably should reverse this logic
}

// CreateInstanceSnapshotRequest is the request body for POST /cloud/instance/snapshot.
type CreateInstanceSnapshotRequest struct {
	Snapshot string `json:"snapshot"`
}

// UpdateInstanceTemplateRequest is the request body for PATCH /cloud/instance/{instance}/template.
type UpdateInstanceTemplateRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// CreateVNCProxyRequest is the request body for POST /cloud/instance/{instance}/vnc.
type CreateVNCProxyRequest struct {
	Zone int `json:"zone"`
}

//
// Security Group
//

// CreateSecurityGroupRequest is the request body for POST /cloud/security-group.
type CreateSecurityGroupRequest struct {
	Template string `json:"template"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *CreateSecurityGroupRequest) SetTemplate(t SecurityGroupTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// CloneSecurityGroupRequest is the request body for POST /cloud/security-group/{sg}/clone.
type CloneSecurityGroupRequest struct {
	Name string `json:"name"`
}

// UpdateSecurityGroupRequest is the request body for PATCH /cloud/security-group/{sg}.
type UpdateSecurityGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// CommitSecurityGroupRequest is the request body for PATCH /cloud/security-group/{sg}/commit.
type CommitSecurityGroupRequest struct {
	All bool `json:"all"`
}

// ChangeSecurityGroupPermissionsRequest is the request body for PATCH /cloud/security-group/{sg}/chmod.
type ChangeSecurityGroupPermissionsRequest struct {
	Permissions hc.Perms `json:"perms"`
}

// ChangeSecurityGroupOwnershipRequest is the request body for PATCH /cloud/security-group/{sg}/chown.
type ChangeSecurityGroupOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameSecurityGroupRequest is the request body for PATCH /cloud/security-group/{sg}/rename.
type RenameSecurityGroupRequest struct {
	Name string `json:"name"`
}

//
// Template
//

// CreateTemplateRequest is the response body for POST /cloud/template.
type CreateTemplateRequest struct {
	Data string `json:"data"`
}

// CloneTemplateRequest is the response body for POST /cloud/template/{template}/clone.
type CloneTemplateRequest struct {
	Name string `json:"name"`
	Disk bool   `json:"disk"`
}

// InstantiateTemplateRequest is the response body for PATCH /cloud/template/{template}/instantiate.
type InstantiateTemplateRequest struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Hold     bool   `json:"hold"`
	DiskCopy bool   `json:"disk_copy"`
}

// UpdateTemplateRequest is the response body for PATCH /cloud/template/{template}.
type UpdateTemplateRequest struct {
	Data  string `json:"data"`
	Merge bool   `json:"merge"`
}

// ChangeTemplatePermissionsRequest is the response body for PATCH /cloud/template/{template}/permissions.
type ChangeTemplatePermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
	Disk        bool     `json:"disk"`
}

// ChangeTemplateOwnershipRequest is the response body for PATCH /cloud/template/{template}/ownership.
type ChangeTemplateOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameTemplateRequest is the response body for PATCH /cloud/template/{template}/name.
type RenameTemplateRequest struct {
	Name string `json:"name"`
}

// LockTemplateRequest is the response body for PATCH /cloud/template/{template}/lock.
type LockTemplateRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

//
// User
//

// CreateUserRequest is the request body for POST /cloud/user.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
	Gids     []int  `json:"gids"`
}

// ChangeUserPasswordRequest is the request body for PATCH /cloud/user/{user}/password.
type ChangeUserPasswordRequest struct {
	Password string `json:"password"`
}

// UserLoginRequest is the request body for POST /cloud/user/login.
type UserLoginRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Duration int    `json:"duration"`
	Group    int    `json:"group"`
}

// UpdateUserRequest is the request body for PATCH /cloud/user/{user}.
type UpdateUserRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// SetTemplate sets r.Template from its structured representation. Either call
// this or set the template string directly.
func (r *UpdateUserRequest) SetTemplate(t UserTemplate) {
	r.Template = mustParseTemplate(t).String()
}

// ChangeUserAuthRequest is the request body for PATCH /cloud/user/{user}/auth.
type ChangeUserAuthRequest struct {
	Driver   string `json:"driver"`
	Password string `json:"password"`
}

// SetUserQuotaRequest is the request body for PATCH /cloud/user/{user}/quota.
type SetUserQuotaRequest struct {
	Template string `json:"template"`
}

// EnableUserRequest is the request body for PATCH /cloud/user/{user}/enable.
type EnableUserRequest struct {
	Enable bool `json:"enable"`
}

// UpdateDefaultUserQuotaRequest is the request body for POST /cloud/user/quota.
type UpdateDefaultUserQuotaRequest struct {
	Template string `json:"template"`
}

//
// VDC - Virtual Data Center
//

// CreateDataCenterRequest is the request body for POST /cloud/datacenter.
type CreateDataCenterRequest struct {
	Template string `json:"template"`
}

// UpdateDataCenterRequest is the request body for PATCH /cloud/datacenter/{datacenter}.
type UpdateDataCenterRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameDataCenterRequest is the request body for PATCH /cloud/datacenter/{datacenter}/name.
type RenameDataCenterRequest struct {
	Name string `json:"name"`
}

//
// VMGroup
//

// CreateInstanceGroupRequest is the request body for POST /cloud/instance-group.
type CreateInstanceGroupRequest struct {
	Template string `json:"template"`
}

// UpdateInstanceGroupRequest is the request body for PATCH /cloud/instance-group/{group}.
type UpdateInstanceGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeInstanceGroupOwnershipRequest is the request body for PATCH /cloud/instance-group/{group}/ownership.
type ChangeInstanceGroupOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// ChangeInstanceGroupPermissionsRequest is the request body for PATCH /cloud/instance-group/{group}/permissions.
type ChangeInstanceGroupPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// RenameInstanceGroupRequest is the request body for PATCH /cloud/instance-group/name.
type RenameInstanceGroupRequest struct {
	Name string `json:"name"`
}

// LockInstanceGroupRequest is the request body for PATCH /cloud/instance-group/lock.
type LockInstanceGroupRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Market
//

// ChangeMarketAppOwnershipRequest is the request body for PATCH /cloud/market/app/{app}/ownership.
type ChangeMarketAppOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// ChangeMarketOwnershipRequest is the request body for PATCH /cloud/market/{market}/ownership.
type ChangeMarketOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// ChangeMarketAppPermissionsRequest is the request body for PATCH /cloud/market/app/{app}/permissions.
type ChangeMarketAppPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeMarketPermissionsRequest is the request body for PATCH /cloud/market/{market}/permissions.
type ChangeMarketPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// CreateMarketRequest is the request body for POST /cloud/market.
type CreateMarketRequest struct {
	Template string `json:"template"`
}

// CreateMarketAppRequest is the request body for POST /cloud/market/{market}/app/{app}.
type CreateMarketAppRequest struct {
	Template string `json:"template"`
}

// EnableMarketAppRequest is the request body for PATCH /cloud/market/app/enable.
type EnableMarketAppRequest struct {
	Enable bool `json:"enable"`
}

// EnableMarketRequest is the request body for PATCH /cloud/market/{market}/enable.
type EnableMarketRequest struct {
	Enable bool `json:"enable"`
}

// LockMarketAppRequest is the request body for PATCH /cloud/market/app/{app}/lock.
type LockMarketAppRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

// UpdateMarketAppRequest is the request body for PATCH /cloud/market/app/{app}.
type UpdateMarketAppRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// UpdateMarketRequest is the request body for PATCH /cloud/market/{market}.
type UpdateMarketRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameMarketAppRequest is the request body for PATCH /cloud/market/app/{app}/name.
type RenameMarketAppRequest struct {
	Name string `json:"name"`
}

// RenameMarketRequest is the request body for PATCH /cloud/market/{market}/name.
type RenameMarketRequest struct {
	Name string `json:"name"`
}

//
// VNet - Virtual Network
//

// CreateNetworkRequest is the request body for POST /cloud/network.
type CreateNetworkRequest struct {
	Template string `json:"template"`
	Cluster  *int   `json:"cluster"`
}

// AddNetworkAddressRangeRequest is the request body for POST /cloud/network/{network}/address-range.
type AddNetworkAddressRangeRequest struct {
	Template string `json:"template"`
}

// UpdateNetworkAddressRangeRequest is the request body for PATCH /cloud/network/{network}/address-range.
type UpdateNetworkAddressRangeRequest struct {
	Template string `json:"template"`
}

// ReserveNetworkRequest is the request body for POST /cloud/network/{network}/reserve.
type ReserveNetworkRequest struct {
	Template string `json:"template"`
}

// HoldNetworkRequest is the request body for PATCH /cloud/network/{network}/hold.
type HoldNetworkRequest struct {
	Template string `json:"template"`
}

// ReleaseNetworkRequest is the request body for PATCH /cloud/network/{network}/release.
type ReleaseNetworkRequest struct {
	Template string `json:"template"`
}

// UpdateNetworkRequest is the request body for PATCH /cloud/network/{network}.
type UpdateNetworkRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeNetworkPermissionsRequest is the request body for PATCH /cloud/network/{network}/permissions.
type ChangeNetworkPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeNetworkOwnershipRequest is the request body for PATCH /cloud/network/{network}/ownership.
type ChangeNetworkOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameNetworkRequest is the request body for PATCH /cloud/network/{network}/name.
type RenameNetworkRequest struct {
	Name string `json:"name"`
}

// LockNetworkRequest is the request body for PATCH /cloud/network/{network}/lock.
type LockNetworkRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

// RecoverNetworkRequest is the request body for PATCH /cloud/network/{network}/recover.
type RecoverNetworkRequest struct {
	Recovery string `json:"recovery"`
}

// CreateNetworkTemplateRequest is the request body for POST /cloud/network/template.
type CreateNetworkTemplateRequest struct {
	Template string `json:"template"`
}

// CloneNetworkTemplateRequest is the request body for POST /cloud/network/template/clone.
type CloneNetworkTemplateRequest struct {
	Name string `json:"name"`
}

// InstantiateNetworkTemplateRequest is the request body for PATCH /cloud/network/template/{template}/instantiate.
type InstantiateNetworkTemplateRequest struct {
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

// UpdateNetworkTemplateRequest is the request body for PATCH /cloud/network/{network}/template.
type UpdateNetworkTemplateRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeNetworkTemplatePermissionsRequest is the request body for PATCH /cloud/network/template/{template}/permissions.
type ChangeNetworkTemplatePermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeNetworkTemplateOwnershipRequest is the request body for PATCH /cloud/network/template/{template}/ownership.
type ChangeNetworkTemplateOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameNetworkTemplateRequest is the request body for PATCH /cloud/network/template/{template}/name.
type RenameNetworkTemplateRequest struct {
	Name string `json:"name"`
}

// LockNetworkTemplateRequest is the request body for PATCH /cloud/network/template/{template}/lock.
type LockNetworkTemplateRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// VRouter
//

// CreateRouterRequest is the request body for POST /cloud/router.
type CreateRouterRequest struct {
	Template string `json:"template"`
}

// InstantiateRouterRequest is the request body for PATCH /cloud/router/{router}/instantiate.
type InstantiateRouterRequest struct {
	Instances        int    `json:"instances"`
	InstanceTemplate int    `json:"instance_template"`
	Name             string `json:"name"`
	Pending          bool   `json:"pending"`
	Extra            string `json:"extra"`
}

// CreateRouterNICRequest is the request body for PATCH /cloud/router/nic.
type CreateRouterNICRequest struct {
	Template string `json:"template"`
}

// UpdateRouterRequest is the request body for PATCH /cloud/router/{router}.
type UpdateRouterRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeRouterPermissionsRequest is the request body for PATCH /cloud/router/{router}/permissions.
type ChangeRouterPermissionsRequest struct {
	Permissions hc.Perms `json:"permissions"`
}

// ChangeRouterOwnershipRequest is the request body for PATCH /cloud/router/{router}/ownership.
type ChangeRouterOwnershipRequest struct {
	User  *int `json:"user"`
	Group *int `json:"group"`
}

// RenameRouterRequest is the request body for PATCH /cloud/router/{router}/name.
type RenameRouterRequest struct {
	Name string `json:"name"`
}

// LockRouterRequest is the request body for PATCH /cloud/router/{router}/lock.
type LockRouterRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Zone
//

// CreateZoneRequest is the request body for POST /cloud/zone.
type CreateZoneRequest struct {
	Template string `json:"template"`
}

// EnableZoneRequest is the request body for PATCH /cloud/zone/{zone}/enable.
type EnableZoneRequest struct {
	Enable bool `json:"enable"`
}

// UpdateZoneRequest is the request body for PATCH /cloud/zone/{zone}.
type UpdateZoneRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameZoneRequest is the request body for PATCH /cloud/zone/{zone}/name.
type RenameZoneRequest struct {
	Name string `json:"name"`
}
