package hc

//
// ACL
//

// CreateACLRequest is the request body for POST /acl.
type CreateACLRequest struct {
	User     int  `json:"user"`
	Resource int  `json:"resource"`
	Rights   int  `json:"rights"`
	Zone     *int `json:"zone,omitempty"`
}

//
// Cluster
//

// CreateClusterRequest is request body for POST /cluster.
type CreateClusterRequest struct {
	Name string `json:"name"`
}

// UpdateClusterRequest is request body for PATCH /cluster/{cluster}.
type UpdateClusterRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameClusterRequest is the request body for PATCH /cluster/{cluster}/name.
type RenameClusterRequest struct {
	Name string `json:"name"`
}

//
// Datastore
//

// CreateDatastoreRequest is the request body for POST /datastore.
type CreateDatastoreRequest struct {
	Template string `json:"template"`
	Cluster  *int   `json:"cluster,omitempty"`
}

// UpdateDatastoreRequest is the request body for PATCH /datastore/{datastore}.
type UpdateDatastoreRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeDatastorePermissionsRequest is the request body for PATCH /datastore/{datastore}/permissions.
type ChangeDatastorePermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeDatastoreOwnershipRequest is the request body for PATCH /datastore/{datastore}/ownership.
type ChangeDatastoreOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameDatastoreRequest is the request body for PATCH /datastore/{datastore}/name.
type RenameDatastoreRequest struct {
	Name string `json:"name"`
}

// EnableDatastoreRequest is the request body for PATCH /datastore/{datastore}/enable.
type EnableDatastoreRequest struct {
	Enable bool `json:"enable"`
}

//
// Document
//

// AllocateDocumentRequest is the request body for POST /document.
type AllocateDocumentRequest struct {
	Template string `json:"template"`
	Type     int    `json:"type"`
}

// CloneDocumentRequest is the request body for POST /document/{document}/clone.
type CloneDocumentRequest struct {
	Name string `json:"name"`
}

// UpdateDocumentRequest is the request body for PATCH /document/{document}.
type UpdateDocumentRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeDocumentPermissionsRequest is the request body for PATCH /document/{document}/permissions.
type ChangeDocumentPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeDocumentOwnershipRequest is the request body for PATCH /document/{document}/ownership.
type ChangeDocumentOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameDocumentRequest is the request body for PATCH /document/{document}/name.
type RenameDocumentRequest struct {
	Name string `json:"name"`
}

// LockDocumentRequest is the request body for PATCH /document/{document}/lock.
type LockDocumentRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Group
//

// CreateGroupRequest is the request body for POST /group.
type CreateGroupRequest struct {
	Name string `json:"name"`
}

// UpdateGroupRequest is the request body for PATCH /group.
type UpdateGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// AddGroupAdminRequest is the request body for POST /group/admin.
type AddGroupAdminRequest struct {
	User int `json:"user"`
}

// SetGroupQuotaRequest is the request body for POST /group/quota.
type SetGroupQuotaRequest struct {
	Template string `json:"template"`
}

// UpdateGroupQuotaRequest is the request body for PATCH /group/{group}/quota.
type UpdateGroupQuotaRequest struct {
	Template string `json:"template"`
}

//
// Hook
//

// CreateHookRequest is the request body for POST /hook.
type CreateHookRequest struct {
	Template string `json:"template"`
}

// UpdateHookRequest is the request body for PATCH /hook/{hook}.
type UpdateHookRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameHookRequest is the request body for PATCH /hook/{hook}/name.
type RenameHookRequest struct {
	Name string `json:"name"`
}

// LockHookRequest is the request body for PATCH /hook/{hook}/lock.
type LockHookRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

// UnlockHookRequest is the request body for PATCH /hook/{hook}/unlock.
type UnlockHookRequest struct {
	Lock int `json:"lock"`
}

// RetryHookRequest is the request body for PATCH /hook/{hook}/retry.
type RetryHookRequest struct {
	Execution int `json:"execution"`
}

//
// Host
//

// CreateHostRequest is the request body for POST /host.
type CreateHostRequest struct {
	Hostname    string `json:"hostname"`
	InfoManager string `json:"info_manager"`
	VMManager   string `json:"vm_manager"`
	Cluster     *int   `json:"cluster,omitempty"`
}

// SetHostStatusRequest is the request body for PATCH /host/{host}/status.
type SetHostStatusRequest struct {
	Status string `json:"status"`
}

// UpdateHostRequest is the request body for PATCH /host/{host}.
type UpdateHostRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameHostRequest is the request body for PATCH /host/{host}/name.
type RenameHostRequest struct {
	Hostname string `json:"hostname"`
}

//
// Image
//

// CreateImageRequest is the request body for POST /image.
type CreateImageRequest struct {
	Template        string `json:"template"`
	Datastore       int    `json:"datastore"`
	EnforceCapacity bool   `json:"enforce_capacity"`
}

// CloneImageRequest is the request body for POST /image/clone.
type CloneImageRequest struct {
	Name      string `json:"name"`
	Datastore *int   `json:"datastore"`
}

// EnableImageRequest is the request body for PATCH /image/{image}/enable.
type EnableImageRequest struct {
	Enable bool `json:"enable"`
}

// SetImagePersistentRequest is the request body for PATCH /image/{image}/persistent.
type SetImagePersistentRequest struct {
	Persistent bool `json:"persistent"`
}

// ChangeImageTypeRequest is the request body for PATCH /image/{image}/type.
type ChangeImageTypeRequest struct {
	Type string `json:"type"`
}

// UpdateImageRequest is the request body for PATCH /image/{image}.
type UpdateImageRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeImagePermissionsRequest is the request body for PATCH /image/{image}/permissions.
type ChangeImagePermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeImageOwnershipRequest is the request body for PATCH /image/{image}/ownership.
type ChangeImageOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameImageRequest is the request body for PATCH /image/{image}/name.
type RenameImageRequest struct {
	Name string `json:"name"`
}

// LockImageRequest is the request body for PATCH /image/{image}/lock.
type LockImageRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Instance
//

// CreateInstanceRequest is the request body for POST /instance.
type CreateInstanceRequest struct {
	Template string `json:"template"`
	Pending  bool   `json:"pending"`
}

// SetInstanceActionRequest is the request body for PATCH /instance/{instance}/action.
type SetInstanceActionRequest struct {
	// Action to perform
	Action string `json:"action" enums:"hold,poweroff,poweroff-hard,reboot,reboot-hard,release,resched,resume,stop,suspend,terminate,terminate-hard,updeploy,updeploy-hard"`
}

// UpdateInstanceConfigRequest is the request body for PATCH /instance/{instance}/config.
type UpdateInstanceConfigRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// DeployInstanceRequest is the request body for POST /instance/{instance}/deploy.
type DeployInstanceRequest struct {
	Host      int    `json:"host"`
	Enforce   bool   `json:"enforce,omitempty"`
	Datastore *int   `json:"datastore,omitempty"`
	Template  string `json:"template"`
}

// CreateInstanceDiskRequest is the request body for POST /instance/disk.
type CreateInstanceDiskRequest struct {
	DiskTemplate string `json:"disk"`
}

// CreateInstanceDiskImageRequest is the request body for POST /instance/disk/image.
type CreateInstanceDiskImageRequest struct {
	Name      string `json:"name"`
	ImageType string `json:"image_type"`
	Snapshot  *int   `json:"snapshot,omitempty"`
}

// ResizeInstanceDiskRequest is the request body for POST /instance/{instance}/disk/{disk}/size.
type ResizeInstanceDiskRequest struct {
	Size int `json:"size"`
}

// CreateInstanceDiskSnapshotRequest is the request body for POST /instance/disk/snapshot.
type CreateInstanceDiskSnapshotRequest struct {
	Description string `json:"description"`
}

// RenameInstanceDiskSnapshotRequest is the request body for PATCH /instance/{instance}/disk/{disk}/snapshot/{snapshot}/name.
type RenameInstanceDiskSnapshotRequest struct {
	Name string `json:"name"`
}

// LockInstanceRequest is the request body for PATCH /instance/{instance}/lock.
type LockInstanceRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test,omitempty"`
}

// MoveInstanceRequest is the request body for PATCH /instance/{instance}/move.
type MoveInstanceRequest struct {
	Host            int    `json:"host"`
	LiveMigrate     bool   `json:"live_migrate"`
	EnforceCapacity bool   `json:"enforce_capacity"` // TODO: probably should reverse this logic
	MigrationType   string `json:"migration_type"`
}

// RenameInstanceRequest is the request body for PATCH /instance/{instance}/name.
type RenameInstanceRequest struct {
	Name string `json:"name"`
}

// CreateInstanceNICRequest is the request body for POST /instance/nic.
type CreateInstanceNICRequest struct {
	NICTemplate string `json:"nic"`
}

// ChangeInstanceOwnershipRequest is the request body for PATCH /instance/{instance}/ownership.
type ChangeInstanceOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// ChangeInstancePermissionsRequest is the request body for PATCH /instance/{instance}/permissions.
type ChangeInstancePermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// RecoverInstanceRequest is the request body for PATCH /instance/{instance}/recover.
type RecoverInstanceRequest struct {
	Operation string `json:"operation"`
}

// AddInstanceScheduleRequest is the request body for POST /instance/schedule.
type AddInstanceScheduleRequest struct {
	Template string `json:"template"`
}

// UpdateInstanceScheduleRequest is the request body for PATCH /instance/{instance}/schedule.
type UpdateInstanceScheduleRequest struct {
	Action   int    `json:"action"`
	Template string `json:"template"`
}

// CalculateInstancesShowbackRequest is the request body for POST /instance/showback.
type CalculateInstancesShowbackRequest struct {
	Month Period `json:"month"`
	Year  Period `json:"year"`
}

// ResizeInstanceRequest is the request body for PATCH /instance/{instance}/size.
type ResizeInstanceRequest struct {
	Template        string `json:"template"`
	EnforceCapacity bool   `json:"enforce_capacity"` // TODO: probably should reverse this logic
}

// CreateInstanceSnapshotRequest is the request body for POST /instance/snapshot.
type CreateInstanceSnapshotRequest struct {
	Snapshot string `json:"snapshot"`
}

// UpdateInstanceTemplateRequest is the request body for PATCH /instance/{instance}/template.
type UpdateInstanceTemplateRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

//
// Security Group
//

// CreateSecurityGroupRequest is the request body for POST /security-group.
type CreateSecurityGroupRequest struct {
	Template string `json:"template"`
}

// CloneSecurityGroupRequest is the request body for POST /security-group/{sg}/clone.
type CloneSecurityGroupRequest struct {
	Name string `json:"name"`
}

// UpdateSecurityGroupRequest is the request body for PATCH /security-group/{sg}.
type UpdateSecurityGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// CommitSecurityGroupRequest is the request body for PATCH /security-group/{sg}/commit.
type CommitSecurityGroupRequest struct {
	All bool `json:"all"`
}

// ChangeSecurityGroupPermissionsRequest is the request body for PATCH /security-group/{sg}/chmod.
type ChangeSecurityGroupPermissionsRequest struct {
	Perms Perms `json:"perms"`
}

// ChangeSecurityGroupOwnershipRequest is the request body for PATCH /security-group/{sg}/chown.
type ChangeSecurityGroupOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameSecurityGroupRequest is the request body for PATCH /security-group/{sg}/rename.
type RenameSecurityGroupRequest struct {
	Name string `json:"name"`
}

//
// Template
//

// CreateTemplateRequest is the response body for POST /template.
type CreateTemplateRequest struct {
	Data string `json:"data"`
}

// CloneTemplateRequest is the response body for POST /template/{template}/clone.
type CloneTemplateRequest struct {
	Name string `json:"name"`
	Disk bool   `json:"disk"`
}

// InstantiateTemplateRequest is the response body for PATCH /template/{template}/instantiate.
type InstantiateTemplateRequest struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Hold     bool   `json:"hold"`
	DiskCopy bool   `json:"disk_copy"`
}

// UpdateTemplateRequest is the response body for PATCH /template/{template}.
type UpdateTemplateRequest struct {
	Data  string `json:"data"`
	Merge bool   `json:"merge"`
}

// ChangeTemplatePermissionsRequest is the response body for PATCH /template/{template}/permissions.
type ChangeTemplatePermissionsRequest struct {
	Permissions Perms `json:"permissions"`
	Disk        bool  `json:"disk"`
}

// ChangeTemplateOwnershipRequest is the response body for PATCH /template/{template}/ownership.
type ChangeTemplateOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameTemplateRequest is the response body for PATCH /template/{template}/name.
type RenameTemplateRequest struct {
	Name string `json:"name"`
}

// LockTemplateRequest is the response body for PATCH /template/{template}/lock.
type LockTemplateRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

//
// User
//

// CreateUserRequest is the request body for POST /user.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
	Gids     []int  `json:"gids"`
}

// ChangeUserPasswordRequest is the request body for PATCH /user/{user}/password.
type ChangeUserPasswordRequest struct {
	Password string `json:"password"`
}

// SetUserLoginRequest is the request body for PATCH /user/login.
type SetUserLoginRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Duration int    `json:"duration"`
	Group    int    `json:"group"`
}

// UpdateUserRequest is the request body for PATCH /user/{user}.
type UpdateUserRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeUserAuthRequest is the request body for PATCH /user/{user}/auth.
type ChangeUserAuthRequest struct {
	Driver   string `json:"driver"`
	Password string `json:"password"`
}

// SetUserQuotaRequest is the request body for PATCH /user/{user}/quota.
type SetUserQuotaRequest struct {
	Template string `json:"template"`
}

// EnableUserRequest is the request body for PATCH /user/{user}/enable.
type EnableUserRequest struct {
	Enable bool `json:"enable"`
}

// UpdateDefaultUserQuotaRequest is the request body for POST /user/quota.
type UpdateDefaultUserQuotaRequest struct {
	Template string `json:"template"`
}

//
// VDC - Virtual Data Center
//

// CreateVDCRequest is the request body for POST /vdc.
type CreateVDCRequest struct {
	Template string `json:"template"`
}

// UpdateVDCRequest is the request body for PATCH /vdc/{vdc}.
type UpdateVDCRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameVDCRequest is the request body for PATCH /vdc/{vdc}/name.
type RenameVDCRequest struct {
	Name string `json:"name"`
}

//
// VMGroup
//

// CreateInstanceGroupRequest is the request body for POST /instance-group.
type CreateInstanceGroupRequest struct {
	Template string `json:"template"`
}

// UpdateInstanceGroupRequest is the request body for PATCH /instance-group/{group}.
type UpdateInstanceGroupRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeInstanceGroupOwnershipRequest is the request body for PATCH /instance-group/{group}/ownership.
type ChangeInstanceGroupOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// ChangeInstanceGroupPermissionsRequest is the request body for PATCH /instance-group/{group}/permissions.
type ChangeInstanceGroupPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// RenameInstanceGroupRequest is the request body for PATCH /instance-group/name.
type RenameInstanceGroupRequest struct {
	Name string `json:"name"`
}

// LockInstanceGroupRequest is the request body for PATCH /instance-group/lock.
type LockInstanceGroupRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Market
//

// ChangeMarketAppOwnershipRequest is the request body for PATCH /market/app/{app}/ownership.
type ChangeMarketAppOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// ChangeMarketOwnershipRequest is the request body for PATCH /market/{market}/ownership.
type ChangeMarketOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// ChangeMarketAppPermissionsRequest is the request body for PATCH /market/app/{app}/permissions.
type ChangeMarketAppPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeMarketPermissionsRequest is the request body for PATCH /market/{market}/permissions.
type ChangeMarketPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// CreateMarketRequest is the request body for POST /market.
type CreateMarketRequest struct {
	Template string `json:"template"`
}

// CreateMarketAppRequest is the request body for POST /market/{market}/app/{app}.
type CreateMarketAppRequest struct {
	Template string `json:"template"`
}

// EnableMarketAppRequest is the request body for PATCH /market/app/enable.
type EnableMarketAppRequest struct {
	Enable bool `json:"enable"`
}

// EnableMarketRequest is the request body for PATCH /market/{market}/enable.
type EnableMarketRequest struct {
	Enable bool `json:"enable"`
}

// LockMarketAppRequest is the request body for PATCH /market/app/{app}/lock.
type LockMarketAppRequest struct {
	Level LockLevel `json:"level"`
	Test  bool      `json:"test"`
}

// UpdateMarketAppRequest is the request body for PATCH /market/app/{app}.
type UpdateMarketAppRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// UpdateMarketRequest is the request body for PATCH /market/{market}.
type UpdateMarketRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameMarketAppRequest is the request body for PATCH /market/app/{app}/name.
type RenameMarketAppRequest struct {
	Name string `json:"name"`
}

// RenameMarketRequest is the request body for PATCH /market/{market}/name.
type RenameMarketRequest struct {
	Name string `json:"name"`
}

//
// VNet - Virtual Network
//

// CreateVNetRequest is the request body for POST /vnet.
type CreateVNetRequest struct {
	Template string `json:"template"`
	Cluster  *int   `json:"cluster,omitempty"`
}

// AddVNetAddressRangeRequest is the request body for POST /vnet/{vnet}/address-range.
type AddVNetAddressRangeRequest struct {
	Template string `json:"template"`
}

// UpdateVNetAddressRangeRequest is the request body for PATCH /vnet/{vnet}/address-range.
type UpdateVNetAddressRangeRequest struct {
	Template string `json:"template"`
}

// ReserveVNetRequest is the request body for POST /vnet/{vnet}/reserve.
type ReserveVNetRequest struct {
	Template string `json:"template"`
}

// HoldVNetRequest is the request body for PATCH /vnet/{vnet}/hold.
type HoldVNetRequest struct {
	Template string `json:"template"`
}

// ReleaseVNetRequest is the request body for PATCH /vnet/{vnet}/release.
type ReleaseVNetRequest struct {
	Template string `json:"template"`
}

// UpdateVNetRequest is the request body for PATCH /vnet/{vnet}.
type UpdateVNetRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeVNetPermissionsRequest is the request body for PATCH /vnet/{vnet}/permissions.
type ChangeVNetPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeVNetOwnershipRequest is the request body for PATCH /vnet/{vnet}/ownership.
type ChangeVNetOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameVNetRequest is the request body for PATCH /vnet/{vnet}/name.
type RenameVNetRequest struct {
	Name string `json:"name"`
}

// LockVNetRequest is the request body for PATCH /vnet/{vnet}/lock.
type LockVNetRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

// RecoverVNetRequest is the request body for PATCH /vnet/{vnet}/recover.
type RecoverVNetRequest struct {
	Recovery string `json:"recovery"`
}

// CreateVNetTemplateRequest is the request body for POST /vnet/template.
type CreateVNetTemplateRequest struct {
	Template string `json:"template"`
}

// CloneVNetTemplateRequest is the request body for POST /vnet/template/clone.
type CloneVNetTemplateRequest struct {
	Name string `json:"name"`
}

// InstantiateVNetTemplateRequest is the request body for PATCH /vnet/template/{template}/instantiate.
type InstantiateVNetTemplateRequest struct {
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

// UpdateVNetTemplateRequest is the request body for PATCH /vnet/{vnet}/template.
type UpdateVNetTemplateRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeVNetTemplatePermissionsRequest is the request body for PATCH /vnet/template/{template}/permissions.
type ChangeVNetTemplatePermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeVNetTemplateOwnershipRequest is the request body for PATCH /vnet/template/{template}/ownership.
type ChangeVNetTemplateOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameVNetTemplateRequest is the request body for PATCH /vnet/template/{template}/name.
type RenameVNetTemplateRequest struct {
	Name string `json:"name"`
}

// LockVNetTemplateRequest is the request body for PATCH /vnet/template/{template}/lock.
type LockVNetTemplateRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// VRouter
//

// CreateVRouterRequest is the request body for POST /vrouter.
type CreateVRouterRequest struct {
	Template string `json:"template"`
}

// InstantiateVRouterRequest is the request body for PATCH /vrouter/{router}/instantiate.
type InstantiateVRouterRequest struct {
	Instances        int    `json:"instances"`
	InstanceTemplate int    `json:"instance_template"`
	Name             string `json:"name"`
	Pending          bool   `json:"pending"`
	Extra            string `json:"extra"`
}

// CreateVRouterNICRequest is the request body for PATCH /vrouter/nic.
type CreateVRouterNICRequest struct {
	Template string `json:"template"`
}

// UpdateVRouterRequest is the request body for PATCH /vrouter/{router}.
type UpdateVRouterRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// ChangeVRouterPermissionsRequest is the request body for PATCH /vrouter/{router}/permissions.
type ChangeVRouterPermissionsRequest struct {
	Permissions Perms `json:"permissions"`
}

// ChangeVRouterOwnershipRequest is the request body for PATCH /vrouter/{router}/ownership.
type ChangeVRouterOwnershipRequest struct {
	User  *int `json:"user,omitempty"`
	Group *int `json:"group,omitempty"`
}

// RenameVRouterRequest is the request body for PATCH /vrouter/{router}/name.
type RenameVRouterRequest struct {
	Name string `json:"name"`
}

// LockVRouterRequest is the request body for PATCH /vrouter/{router}/lock.
type LockVRouterRequest struct {
	Level string `json:"level"`
	Test  bool   `json:"test"`
}

//
// Zone
//

// CreateZoneRequest is the request body for POST /zone.
type CreateZoneRequest struct {
	Template string `json:"template"`
}

// EnableZoneRequest is the request body for PATCH /zone/{zone}/enable.
type EnableZoneRequest struct {
	Enable bool `json:"enable"`
}

// UpdateZoneRequest is the request body for PATCH /zone/{zone}.
type UpdateZoneRequest struct {
	Template string `json:"template"`
	Merge    bool   `json:"merge"`
}

// RenameZoneRequest is the request body for PATCH /zone/{zone}/name.
type RenameZoneRequest struct {
	Name string `json:"name"`
}
