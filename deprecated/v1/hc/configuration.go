package hc

// ConfigAuthMAD is the API payload based on the legacy xmlrpc backend.
type ConfigAuthMAD struct {
	Authn      string `json:"authn" yaml:"authn"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigAuthMADConf is the API payload based on the legacy xmlrpc backend.
type ConfigAuthMADConf struct {
	DriverManagedGroups     string `json:"driver_managed_groups" yaml:"driver_managed_groups"`
	DriverManagedGroupAdmin string `json:"driver_managed_group_admin" yaml:"driver_managed_group_admin"`
	MaxTokenTime            int    `json:"max_token_time" yaml:"max_token_time"`
	Name                    string `json:"name" yaml:"name"`
	PasswordChange          string `json:"password_change" yaml:"password_change"`
}

// ConfigDatastoreMAD is the API payload based on the legacy xmlrpc backend.
type ConfigDatastoreMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigDB is the API payload based on the legacy xmlrpc backend.
type ConfigDB struct {
	Backend       ConfigBackend       `json:"backend" yaml:"backend"`
	CompareBinary ConfigCompareBinary `json:"compare_binary" yaml:"compare_binary"`
	Connections   int                 `json:"connections" yaml:"connections"`
	DBName        string              `json:"dbname" yaml:"dbname"`
	Passwd        string              `json:"passwd" yaml:"passwd"`
	Port          int                 `json:"port" yaml:"port"`
	Server        string              `json:"server" yaml:"server"`
	User          string              `json:"user" yaml:"user"`
	Timeout       int                 `json:"timeout" yaml:"timeout"`
}

// ConfigDefaultCost is the API payload based on the legacy xmlrpc backend.
type ConfigDefaultCost struct {
	CPUCost    int `json:"cpucost" yaml:"cpucost"`
	DiskCost   int `json:"disk_cost" yaml:"disk_cost"`
	MemoryCost int `json:"memory_cost" yaml:"memory_cost"`
}

// ConfigDSMADConf is the API payload based on the legacy xmlrpc backend.
type ConfigDSMADConf struct {
	MarketplaceActions string               `json:"marketplace_actions" yaml:"marketplace_actions"`
	Name               string               `json:"name" yaml:"name"`
	PersistentOnly     ConfigPersistentOnly `json:"persistent_only" yaml:"persistent_only"`
	RequiredAttrs      string               `json:"required_attrs" yaml:"required_attrs"`
}

// ConfigFederation is the API payload based on the legacy xmlrpc backend.
type ConfigFederation struct {
	MasterOned string     `json:"master_oned" yaml:"master_oned"`
	Mode       ConfigMode `json:"mode" yaml:"mode"`
	ServerID   int        `json:"server_id" yaml:"server_id"`
	ZoneID     int        `json:"zone_id" yaml:"zone_id"`
}

// ConfigHmMAD is the API payload based on the legacy xmlrpc backend.
type ConfigHmMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigHookLogConf is the API payload based on the legacy xmlrpc backend.
type ConfigHookLogConf struct {
	LogRetention int `json:"log_retention" yaml:"log_retention"`
}

// ConfigImMAD is the API payload based on the legacy xmlrpc backend.
type ConfigImMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
	Name       string `json:"name" yaml:"name"`
	Threads    int    `json:"threads" yaml:"threads"`
}

// ConfigIpamMAD is the API payload based on the legacy xmlrpc backend.
type ConfigIpamMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigLog is the API payload based on the legacy xmlrpc backend.
type ConfigLog struct {
	DebugLevel           int    `json:"debug_level" yaml:"debug_level"`
	System               string `json:"system" yaml:"system"`
	UseInstancesLocation string `json:"use_instances_location" yaml:"use_instances_location"`
}

// ConfigMarketMAD is the API payload based on the legacy xmlrpc backend.
type ConfigMarketMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigMarketMADConf is the API payload based on the legacy xmlrpc backend.
type ConfigMarketMADConf struct {
	AppActions    string       `json:"app_actions" yaml:"app_actions"`
	Name          string       `json:"name" yaml:"name"`
	Public        ConfigPublic `json:"public" yaml:"public"`
	RequiredAttrs string       `json:"required_attrs" yaml:"required_attrs"`
	SunstoneName  string       `json:"sunstone_name" yaml:"sunstone_name"`
}

// Configuration is the API payload based on the legacy xmlrpc backend.
type Configuration struct {
	APIListOrder                     []string                     `json:"api_list_order" yaml:"api_list_order"`
	AuthMAD                          []ConfigAuthMAD              `json:"auth_mad" yaml:"auth_mad"`
	AuthMADConf                      []ConfigAuthMADConf          `json:"auth_madconf" yaml:"auth_madconf"`
	ClusterEncryptedAttr             []string                     `json:"cluster_encrypted_attr" yaml:"cluster_encrypted_attr"`
	DatastoreCapacityCheck           []string                     `json:"datastore_capacity_check" yaml:"datastore_capacity_check"`
	DatastoreEncryptedAttr           []string                     `json:"datastore_encrypted_attr" yaml:"datastore_encrypted_attr"`
	DatastoreLocation                []string                     `json:"datastore_location" yaml:"datastore_location"`
	DatastoreMAD                     []ConfigDatastoreMAD         `json:"datastore_mad" yaml:"datastore_mad"`
	DB                               ConfigDB                     `json:"db" yaml:"db"`
	DefaultAuth                      []string                     `json:"default_auth" yaml:"default_auth"`
	DefaultCdromDevicePrefix         []string                     `json:"default_cdrom_device_prefix" yaml:"default_cdrom_device_prefix"`
	DefaultCost                      []ConfigDefaultCost          `json:"default_cost" yaml:"default_cost"`
	DefaultDevicePrefix              []string                     `json:"default_device_prefix" yaml:"default_device_prefix"`
	DefaultImagePersistent           []string                     `json:"default_image_persistent" yaml:"default_image_persistent"`
	DefaultImagePersistentNew        []string                     `json:"default_image_persistent_new" yaml:"default_image_persistent_new"`
	DefaultImageType                 []string                     `json:"default_image_type" yaml:"default_image_type"`
	DefaultUmask                     []string                     `json:"default_umask" yaml:"default_umask"`
	DefaultVDCClusterDatastoreACL    []string                     `json:"default_vdccluster_datastore_acl" yaml:"default_vdccluster_datastore_acl"`
	DefaultVDCClusterHostACL         []string                     `json:"default_vdccluster_host_acl" yaml:"default_vdccluster_host_acl"`
	DefaultVDCClusterNetACL          []string                     `json:"default_vdccluster_net_acl" yaml:"default_vdccluster_net_acl"`
	DefaultVDCDatastoreACL           []string                     `json:"default_vdcdatastore_acl" yaml:"default_vdcdatastore_acl"`
	DefaultVDCHostACL                []string                     `json:"default_vdchost_acl" yaml:"default_vdchost_acl"`
	DefaultVDCVnetACL                []string                     `json:"default_vdcvnet_acl" yaml:"default_vdcvnet_acl"`
	DocumentEncryptedAttr            []string                     `json:"document_encrypted_attr" yaml:"document_encrypted_attr"`
	DSMADConf                        []ConfigDSMADConf            `json:"dsmadconf" yaml:"dsmadconf"`
	DSMonitorInstanceDisk            int                          `json:"dsmonitor_instance_disk" yaml:"dsmonitor_instance_disk"`
	EnableOtherPermissions           ConfigEnableOtherPermissions `json:"enable_other_permissions" yaml:"enable_other_permissions"`
	Federation                       ConfigFederation             `json:"federation" yaml:"federation"`
	GroupRestrictedAttr              []string                     `json:"group_restricted_attr" yaml:"group_restricted_attr"`
	HmMAD                            ConfigHmMAD                  `json:"hm_mad" yaml:"hm_mad"`
	HookLogConf                      ConfigHookLogConf            `json:"hook_log_conf" yaml:"hook_log_conf"`
	HostEncryptedAttr                []string                     `json:"host_encrypted_attr" yaml:"host_encrypted_attr"`
	ImageEncryptedAttr               []string                     `json:"image_encrypted_attr" yaml:"image_encrypted_attr"`
	ImageRestrictedAttr              []string                     `json:"image_restricted_attr" yaml:"image_restricted_attr"`
	ImMAD                            []ConfigImMAD                `json:"im_mad" yaml:"im_mad"`
	InheritDatastoreAttr             []string                     `json:"inherit_datastore_attr" yaml:"inherit_datastore_attr"`
	InheritImageAttr                 []string                     `json:"inherit_image_attr" yaml:"inherit_image_attr"`
	InheritVnetAttr                  []string                     `json:"inherit_vnet_attr" yaml:"inherit_vnet_attr"`
	IpamMAD                          []ConfigIpamMAD              `json:"ipam_mad" yaml:"ipam_mad"`
	KeepaliveMaxConn                 []int                        `json:"keepalive_max_conn" yaml:"keepalive_max_conn"`
	KeepaliveTimeout                 []int                        `json:"keepalive_timeout" yaml:"keepalive_timeout"`
	ListenAddress                    []string                     `json:"listen_address" yaml:"listen_address"`
	Log                              []ConfigLog                  `json:"log" yaml:"log"`
	LogCallFormat                    []string                     `json:"log_call_format" yaml:"log_call_format"`
	MACPrefix                        []string                     `json:"macprefix" yaml:"macprefix"`
	ManagerTimer                     []int                        `json:"manager_timer" yaml:"manager_timer"`
	MarketMAD                        []ConfigMarketMAD            `json:"market_mad" yaml:"market_mad"`
	MarketMADConf                    []ConfigMarketMADConf        `json:"market_madconf" yaml:"market_madconf"`
	MaxConn                          int                          `json:"max_conn" yaml:"max_conn"`
	MaxConnBacklog                   int                          `json:"max_conn_backlog" yaml:"max_conn_backlog"`
	MessageSize                      int                          `json:"message_size" yaml:"message_size"`
	MonitoringIntervalDatastore      int                          `json:"monitoring_interval_datastore" yaml:"monitoring_interval_datastore"`
	MonitoringIntervalDBUpdate       int                          `json:"monitoring_interval_dbupdate" yaml:"monitoring_interval_dbupdate"`
	MonitoringIntervalHost           int                          `json:"monitoring_interval_host" yaml:"monitoring_interval_host"`
	MonitoringIntervalMarket         int                          `json:"monitoring_interval_market" yaml:"monitoring_interval_market"`
	MonitoringIntervalInstance       int                          `json:"monitoring_interval_instance" yaml:"monitoring_interval_instance"`
	NetworkSize                      int                          `json:"network_size" yaml:"network_size"`
	Key                              []string                     `json:"one_key" yaml:"one_key"`
	PCIPassthroughBus                string                       `json:"pcipassthrough_bus" yaml:"pcipassthrough_bus"`
	Port                             int                          `json:"port" yaml:"port"`
	Raft                             ConfigRaft                   `json:"raft" yaml:"raft"`
	RPCLog                           string                       `json:"rpclog" yaml:"rpclog"`
	ScriptsRemoteDir                 string                       `json:"scripts_remote_dir" yaml:"scripts_remote_dir"`
	SessionExpirationTime            int                          `json:"session_expiration_time" yaml:"session_expiration_time"`
	ShowbackOnlyRunning              string                       `json:"showback_only_running" yaml:"showback_only_running"`
	Timeout                          int                          `json:"timeout" yaml:"timeout"`
	TmMAD                            []ConfigTmMAD                `json:"tm_mad" yaml:"tm_mad"`
	TmMADConf                        []ConfigTmMADConf            `json:"tm_madconf" yaml:"tm_madconf"`
	UserEncryptedAttr                []string                     `json:"user_encrypted_attr" yaml:"user_encrypted_attr"`
	UserRestrictedAttr               []string                     `json:"user_restricted_attr" yaml:"user_restricted_attr"`
	VLANIDs                          ConfigVLANIDs                `json:"vlanids" yaml:"vlanids"`
	InstanceAdminOperations          string                       `json:"instance_admin_operations" yaml:"instance_admin_operations"`
	InstanceEncryptedAttr            []string                     `json:"instance_encrypted_attr" yaml:"instance_encrypted_attr"`
	InstanceMAD                      []ConfigInstanceMAD          `json:"instance_mad" yaml:"instance_mad"`
	InstanceManageOperations         string                       `json:"instance_manage_operations" yaml:"instance_manage_operations"`
	InstanceMonitoringExpirationTime []string                     `json:"instance_monitoring_expiration_time" yaml:"instance_monitoring_expiration_time"`
	InstanceRestrictedAttr           []string                     `json:"instance_restricted_attr" yaml:"instance_restricted_attr"`
	InstanceSnapshotFactor           string                       `json:"instance_snapshot_factor" yaml:"instance_snapshot_factor"`
	InstanceSubmitOnHold             ConfigInstanceSubmitOnHold   `json:"instance_submit_on_hold" yaml:"instance_submit_on_hold"`
	InstanceUseOperations            []string                     `json:"instance_use_operations" yaml:"instance_use_operations"`
	VNCPorts                         ConfigVNCPorts               `json:"vncports" yaml:"vncports"`
	VNetEncryptedAttr                []string                     `json:"vnet_encrypted_attr" yaml:"vnet_encrypted_attr"`
	VNetRestrictedAttr               []string                     `json:"vnet_restricted_attr" yaml:"vnet_restricted_attr"`
	VNMADConf                        []ConfigVNMADConf            `json:"vnmadconf" yaml:"vnmadconf"`
	VxlanIDs                         ConfigVxlanIDs               `json:"vxlan_ids" yaml:"vxlan_ids"`
}

// ConfigRaft is the API payload based on the legacy xmlrpc backend.
type ConfigRaft struct {
	BroadcastTimeoutMS int `json:"broadcast_timeout_ms" yaml:"broadcast_timeout_ms"`
	ElectionTimeoutMS  int `json:"election_timeout_ms" yaml:"election_timeout_ms"`
	LimitPurge         int `json:"limit_purge" yaml:"limit_purge"`
	LogPurgeTimeout    int `json:"log_purge_timeout" yaml:"log_purge_timeout"`
	LogRetention       int `json:"log_retention" yaml:"log_retention"`
	XmlrpcTimeoutMS    int `json:"xmlrpc_timeout_ms" yaml:"xmlrpc_timeout_ms"`
}

// ConfigTmMAD is the API payload based on the legacy xmlrpc backend.
type ConfigTmMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// ConfigTmMADConf is the API payload based on the legacy xmlrpc backend.
type ConfigTmMADConf struct {
	AllowOrphans      string `json:"allow_orphans" yaml:"allow_orphans"`
	CloneTarget       string `json:"clone_target" yaml:"clone_target"`
	CloneTargetShared string `json:"clone_target_shared" yaml:"clone_target_shared"`
	CloneTargetSSH    string `json:"clone_target_ssh" yaml:"clone_target_ssh"`
	DiskTypeShared    string `json:"disk_type_shared" yaml:"disk_type_shared"`
	DiskTypeSSH       string `json:"disk_type_ssh" yaml:"disk_type_ssh"`
	Driver            string `json:"driver" yaml:"driver"`
	DSMigrate         string `json:"dsmigrate" yaml:"dsmigrate"`
	LnTarget          string `json:"ln_target" yaml:"ln_target"`
	LnTargetShared    string `json:"ln_target_shared" yaml:"ln_target_shared"`
	LnTargetSSH       string `json:"ln_target_ssh" yaml:"ln_target_ssh"`
	Name              string `json:"name" yaml:"name"`
	Shared            string `json:"shared" yaml:"shared"`
	TmMADSystem       string `json:"tm_madsystem" yaml:"tm_madsystem"`
}

// ConfigVLANIDs is the API payload based on the legacy xmlrpc backend.
type ConfigVLANIDs struct {
	Reserved string `json:"reserved" yaml:"reserved"`
	Start    int    `json:"start" yaml:"start"`
}

// ConfigInstanceMAD is the API payload based on the legacy xmlrpc backend.
type ConfigInstanceMAD struct {
	Arguments                string                `json:"arguments" yaml:"arguments"`
	Default                  string                `json:"default" yaml:"default"`
	Executable               string                `json:"executable" yaml:"executable"`
	ImportedInstancesActions string                `json:"imported_instances_actions" yaml:"imported_instances_actions"`
	Name                     string                `json:"name" yaml:"name"`
	SunstoneName             string                `json:"sunstone_name" yaml:"sunstone_name"`
	Type                     string                `json:"type" yaml:"type"`
	KeepSnapshots            ConfigKeepSnapshots   `json:"keep_snapshots" yaml:"keep_snapshots"`
	ColdNICAttach            ConfigColdNICAttach   `json:"cold_nicattach" yaml:"cold_nicattach"`
	DSLiveMigration          ConfigDSLiveMigration `json:"dslive_migration" yaml:"dslive_migration"`
	LiveResize               ConfigLiveResize      `json:"live_resize" yaml:"live_resize"`
}

// ConfigVNCPorts is the API payload based on the legacy xmlrpc backend.
type ConfigVNCPorts struct {
	Reserved string `json:"reserved" yaml:"reserved"`
	Start    int    `json:"start" yaml:"start"`
}

// ConfigVNMADConf is the API payload based on the legacy xmlrpc backend.
type ConfigVNMADConf struct {
	BridgeType string `json:"bridge_type" yaml:"bridge_type"`
	Name       string `json:"name" yaml:"name"`
}

// ConfigVxlanIDs is the API payload based on the legacy xmlrpc backend.
type ConfigVxlanIDs struct {
	Start int `json:"start" yaml:"start"`
}

// ConfigBackend is a type alias for string.
type ConfigBackend string

// ConfigColdNICAttach is a type alias for string.
type ConfigColdNICAttach string

// ConfigCompareBinary is a type alias for string.
type ConfigCompareBinary string

// ConfigDSLiveMigration is a type alias for string.
type ConfigDSLiveMigration string

// ConfigEnableOtherPermissions is a type alias for string.
type ConfigEnableOtherPermissions string

// ConfigKeepSnapshots is a type alias for string.
type ConfigKeepSnapshots string

// ConfigLiveResize is a type alias for string.
type ConfigLiveResize string

// ConfigMode is a type alias for string.
type ConfigMode string

// ConfigPersistentOnly is a type alias for string.
type ConfigPersistentOnly string

// ConfigPublic is a type alias for string.
type ConfigPublic string

// ConfigInstanceSubmitOnHold is a type alias for string.
type ConfigInstanceSubmitOnHold string
