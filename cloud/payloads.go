package cloud

import (
	"github.com/softiron/manifold-api/cloud/config"
	"github.com/softiron/manifold-api/cloud/instance"
	"github.com/softiron/manifold-api/internal/api"
)

// AcctHistory is the API payload based on the legacy xmlrpc backend.
type AcctHistory struct {
	instance.History
	Instance Instance `json:"instance" yaml:"instance"`
}

// ACL is the API payload based on the legacy xmlrpc backend.
type ACL struct {
	ID int `json:"id" yaml:"id"`
	//           32 bits                 32 bits
	//  +-----------------------+-----------------------+
	//  | Type (user,group,all) | user/group ID         |
	//  +-----------------------+-----------------------+
	User int64 `json:"user" yaml:"user"`
	//           32 bits                 32 bits
	//  +-----------------------+-----------------------+
	//  | Type (VM, Host...)    | resource ID           |
	//  +-----------------------+-----------------------+
	Resource int64 `json:"resource" yaml:"resource"`
	//                      64 bits
	//  +-----------------------------------------------+
	//  | Actions (MANAGE, CREATE, USE...               |
	//  +-----------------------------------------------+
	Rights int64 `json:"rights" yaml:"rights"`
	//           32 bits                 32 bits
	//  +-----------------------+-----------------------+
	//  | Type (individual,all) | zone ID               |
	//  +-----------------------+-----------------------+
	Zone   int64  `json:"zone" yaml:"zone"`
	String string `json:"string" yaml:"string"`
}

// Cluster is the API payload based on the legacy xmlrpc backend.
type Cluster struct {
	ID         int      `json:"id" yaml:"id"`
	Name       string   `json:"name" yaml:"name"`
	Hosts      []int    `json:"hosts" yaml:"hosts"`
	Datastores []int    `json:"datastores" yaml:"datastores"`
	Networks   []int    `json:"networks" yaml:"networks"`
	Template   Template `json:"template" yaml:"template"`
}

// ClusterTemplate is the API payload based on the legacy xmlrpc backend.
type ClusterTemplate struct {
	ReservedCPU    string
	ReservedMemory string
}

// Document is the API payload based on the legacy xmlrpc backend.
type Document struct {
	ID          int         `json:"id" yaml:"id"`
	UserID      int         `json:"user_id" yaml:"user_id"`
	GroupID     int         `json:"group_id" yaml:"group_id"`
	UserName    string      `json:"user_name" yaml:"user_name"`
	GroupName   string      `json:"group_name" yaml:"group_name"`
	Name        string      `json:"name" yaml:"name"`
	Type        string      `json:"type" yaml:"type"`
	Permissions Permissions `json:"permissions" yaml:"permissions"`
	Lock        Lock        `json:"lock" yaml:"lock"`
	Template    Template    `json:"template" yaml:"template"`
}

// Instance is the API payload based on the legacy xmlrpc backend.
type Instance struct {
	ID             int                      `json:"id" yaml:"id"`
	UserID         int                      `json:"user_id" yaml:"user_id"`
	GroupID        int                      `json:"group_id" yaml:"group_id"`
	UserName       string                   `json:"user_name" yaml:"user_name"`
	GroupName      string                   `json:"group_name" yaml:"group_name"`
	Name           string                   `json:"name" yaml:"name"`
	Permissions    Permissions              `json:"permissions" yaml:"permissions"`
	LastPoll       api.Time                 `json:"last_poll" yaml:"last_poll"`
	State          instance.State           `json:"state" yaml:"state"`
	LCMState       LCMState                 `json:"lcm_state" yaml:"lcm_state"`
	PrevState      instance.State           `json:"prev_state" yaml:"prev_state"`
	PrevLCMState   LCMState                 `json:"prev_lcm_state" yaml:"prev_lcm_state"`
	Reschedule     bool                     `json:"reschedule" yaml:"reschedule"`
	StartTime      api.Time                 `json:"start_time" yaml:"start_time"`
	EndTime        api.Time                 `json:"end_time" yaml:"end_time"`
	DeployID       string                   `json:"deploy_id" yaml:"deploy_id"`
	Monitoring     instance.Monitoring      `json:"monitoring" yaml:"monitoring"`
	Template       Template                 `json:"template" yaml:"template"`
	UserTemplate   Template                 `json:"user_template" yaml:"user_template"`
	HistoryRecords []instance.History       `json:"history_records" yaml:"history_records"`
	Snapshots      []instance.DiskSnapshots `json:"snapshots" yaml:"snapshots"`
}

// LockedInstance is an Instance with a Lock.
type LockedInstance struct {
	Instance
	Lock Lock `json:"lock" yaml:"lock"`
}

// Marketplace is the API payload based on the legacy xmlrpc backend.
type Marketplace struct {
	ID              int         `json:"id" yaml:"id"`
	UID             int         `json:"uid" yaml:"uid"`
	GID             int         `json:"gid" yaml:"gid"`
	UserName        string      `json:"user_name" yaml:"user_name"`
	GroupName       string      `json:"group_name" yaml:"group_name"`
	Name            string      `json:"name" yaml:"name"`
	State           int         `json:"state" yaml:"state"`
	MarketMAD       string      `json:"market_mad" yaml:"market_mad"`
	ZoneID          string      `json:"zone_id" yaml:"zone_id"`
	TotalMB         int         `json:"total_mb" yaml:"total_mb"`
	FreeMB          int         `json:"free_mb" yaml:"free_mb"`
	UsedMB          int         `json:"used_mb" yaml:"used_mb"`
	MarketplaceApps []int       `json:"marketplace_apps" yaml:"marketplace_apps"`
	Permissions     Permissions `json:"permissions" yaml:"permissions"`
	Template        Template    `json:"template" yaml:"template"`
}

// MarketplaceApp is the API payload based on the legacy xmlrpc backend.
type MarketplaceApp struct {
	ID            int         `json:"id" yaml:"id"`
	UID           int         `json:"uid" yaml:"uid"`
	GID           int         `json:"gid" yaml:"gid"`
	UserName      string      `json:"user_name" yaml:"user_name"`
	GroupName     string      `json:"group_name" yaml:"group_name"`
	Lock          Lock        `json:"lock" yaml:"lock"`
	Regtime       int         `json:"regtime" yaml:"regtime"`
	Name          string      `json:"name" yaml:"name"`
	ZoneID        string      `json:"zone_id" yaml:"zone_id"`
	OriginID      string      `json:"origin_id" yaml:"origin_id"`
	Source        string      `json:"source" yaml:"source"`
	MD5           string      `json:"md5" yaml:"md5"`
	Size          int         `json:"size" yaml:"size"`
	Description   string      `json:"description" yaml:"description"`
	Version       string      `json:"version" yaml:"version"`
	Format        string      `json:"format" yaml:"format"`
	AppTemplate64 string      `json:"apptemplate64" yaml:"apptemplate64"`
	MarketplaceID int         `json:"marketplace_id" yaml:"marketplace_id"`
	Marketplace   string      `json:"marketplace" yaml:"marketplace"`
	State         int         `json:"state" yaml:"state"`
	Type          int         `json:"type" yaml:"type"`
	Permissions   Permissions `json:"permissions" yaml:"permissions"`
	Template      Template    `json:"template" yaml:"template"`
}

// RaftStatus is the API payload based on the legacy xmlrpc backend.
type RaftStatus struct {
	ServerID    int `json:"server_id" yaml:"server_id"`
	State       int `json:"state" yaml:"state"`
	Term        int `json:"term" yaml:"term"`
	VotedFor    int `json:"voted_for" yaml:"voted_for"`
	Commit      int `json:"commit" yaml:"commit"`
	LogIndex    int `json:"log_index" yaml:"log_index"`
	LogTerm     int `json:"log_term" yaml:"log_term"`
	FedLogIndex int `json:"fed_log_index" yaml:"fed_log_index"`
}

// Showback is the API payload based on the legacy xmlrpc backend.
type Showback struct {
	InstanceID   int     `json:"instance" yaml:"instance"`
	InstanceName string  `json:"instance_name" yaml:"instance_name"`
	UID          int     `json:"uid" yaml:"uid"`
	GID          int     `json:"gid" yaml:"gid"`
	UserName     string  `json:"user_name" yaml:"user_name"`
	GroupName    string  `json:"group_name" yaml:"group_name"`
	Year         int     `json:"year" yaml:"year"`
	Month        int     `json:"month" yaml:"month"`
	CPUCost      float32 `json:"cpucost" yaml:"cpucost"`
	MemoryCost   float32 `json:"memory_cost" yaml:"memory_cost"`
	DiskCost     float32 `json:"disk_cost" yaml:"disk_cost"`
	TotalCost    float32 `json:"total_cost" yaml:"total_cost"`
	Hours        float32 `json:"hours" yaml:"hours"`
	Rhours       float32 `json:"rhours" yaml:"rhours"`
}

// InstanceGroup is the API payload based on the legacy xmlrpc backend.
type InstanceGroup struct {
	ID          int                  `json:"id" yaml:"id"`
	UID         int                  `json:"uid" yaml:"uid"`
	GID         int                  `json:"gid" yaml:"gid"`
	UserName    string               `json:"user_name" yaml:"user_name"`
	GroupName   string               `json:"group_name" yaml:"group_name"`
	Name        string               `json:"name" yaml:"name"`
	Permissions Permissions          `json:"permissions" yaml:"permissions"`
	Lock        Lock                 `json:"lock" yaml:"lock"`
	Roles       []instance.GroupRole `json:"roles" yaml:"roles"`
	Template    Template             `json:"template" yaml:"template"`
}

// InstanceTemplate is the API payload based on the legacy xmlrpc backend.
type InstanceTemplate struct {
	ID          int         `json:"id" yaml:"id"`
	UID         int         `json:"uid" yaml:"uid"`
	GID         int         `json:"gid" yaml:"gid"`
	Uname       string      `json:"uname" yaml:"uname"`
	Gname       string      `json:"gname" yaml:"gname"`
	Name        string      `json:"name" yaml:"name"`
	Lock        Lock        `json:"lock" yaml:"lock"`
	Permissions Permissions `json:"permissions" yaml:"permissions"`
	Regtime     int         `json:"regtime" yaml:"regtime"`
	Template    Template    `json:"template" yaml:"template"`
}

// NetworkTemplate is the API payload based on the legacy xmlrpc backend.
type NetworkTemplate struct {
	ID          int         `json:"id" yaml:"id"`
	UID         int         `json:"uid" yaml:"uid"`
	GID         int         `json:"gid" yaml:"gid"`
	UserName    string      `json:"user_name" yaml:"user_name"`
	GroupName   string      `json:"group_name" yaml:"group_name"`
	Name        string      `json:"name" yaml:"name"`
	Lock        Lock        `json:"lock" yaml:"lock"`
	Permissions Permissions `json:"permissions" yaml:"permissions"`
	Regtime     int         `json:"regtime" yaml:"regtime"`
	Template    Template    `json:"template" yaml:"template"`
}

// Router is the API payload based on the legacy xmlrpc backend.
type Router struct {
	ID          int         `json:"id" yaml:"id"`
	UID         int         `json:"uid" yaml:"uid"`
	GID         int         `json:"gid" yaml:"gid"`
	UserName    string      `json:"user_name" yaml:"user_name"`
	GroupName   string      `json:"group_name" yaml:"group_name"`
	Name        string      `json:"name" yaml:"name"`
	Permissions Permissions `json:"permissions" yaml:"permissions"`
	Lock        Lock        `json:"lock" yaml:"lock"`
	Instances   []int       `json:"instances" yaml:"instances"`
	Template    Template    `json:"template" yaml:"template"`
}

// Configuration is the API payload based on the legacy xmlrpc backend.
type Configuration struct {
	APIListOrder                     []string               `json:"api_list_order" yaml:"api_list_order"`
	AuthMAD                          []config.AuthMAD       `json:"auth_mad" yaml:"auth_mad"`
	AuthMADConf                      []config.AuthMADConf   `json:"auth_madconf" yaml:"auth_madconf"`
	ClusterEncryptedAttr             []string               `json:"cluster_encrypted_attr" yaml:"cluster_encrypted_attr"`
	DatastoreCapacityCheck           []string               `json:"datastore_capacity_check" yaml:"datastore_capacity_check"`
	DatastoreEncryptedAttr           []string               `json:"datastore_encrypted_attr" yaml:"datastore_encrypted_attr"`
	DatastoreLocation                []string               `json:"datastore_location" yaml:"datastore_location"`
	DatastoreMAD                     []config.DatastoreMAD  `json:"datastore_mad" yaml:"datastore_mad"`
	DB                               config.DB              `json:"db" yaml:"db"`
	DefaultAuth                      []string               `json:"default_auth" yaml:"default_auth"`
	DefaultCdromDevicePrefix         []string               `json:"default_cdrom_device_prefix" yaml:"default_cdrom_device_prefix"`
	DefaultCost                      []config.DefaultCost   `json:"default_cost" yaml:"default_cost"`
	DefaultDevicePrefix              []string               `json:"default_device_prefix" yaml:"default_device_prefix"`
	DefaultImagePersistent           []string               `json:"default_image_persistent" yaml:"default_image_persistent"`
	DefaultImagePersistentNew        []string               `json:"default_image_persistent_new" yaml:"default_image_persistent_new"`
	DefaultImageType                 []string               `json:"default_image_type" yaml:"default_image_type"`
	DefaultUmask                     []string               `json:"default_umask" yaml:"default_umask"`
	DefaultVDCClusterDatastoreACL    []string               `json:"default_vdccluster_datastore_acl" yaml:"default_vdccluster_datastore_acl"`
	DefaultVDCClusterHostACL         []string               `json:"default_vdccluster_host_acl" yaml:"default_vdccluster_host_acl"`
	DefaultVDCClusterNetACL          []string               `json:"default_vdccluster_net_acl" yaml:"default_vdccluster_net_acl"`
	DefaultVDCDatastoreACL           []string               `json:"default_vdcdatastore_acl" yaml:"default_vdcdatastore_acl"`
	DefaultVDCHostACL                []string               `json:"default_vdchost_acl" yaml:"default_vdchost_acl"`
	DefaultVDCVnetACL                []string               `json:"default_vdcvnet_acl" yaml:"default_vdcvnet_acl"`
	DocumentEncryptedAttr            []string               `json:"document_encrypted_attr" yaml:"document_encrypted_attr"`
	DSMADConf                        []config.DSMADConf     `json:"dsmadconf" yaml:"dsmadconf"`
	DSMonitorInstanceDisk            int                    `json:"dsmonitor_instance_disk" yaml:"dsmonitor_instance_disk"`
	EnableOtherPermissions           string                 `json:"enable_other_permissions" yaml:"enable_other_permissions"`
	Federation                       config.Federation      `json:"federation" yaml:"federation"`
	GroupRestrictedAttr              []string               `json:"group_restricted_attr" yaml:"group_restricted_attr"`
	HookManagerMAD                   config.HookManagerMAD  `json:"hook_manager_mad" yaml:"hook_manager_mad"`
	HookLogConf                      config.HookLogConf     `json:"hook_log_conf" yaml:"hook_log_conf"`
	HostEncryptedAttr                []string               `json:"host_encrypted_attr" yaml:"host_encrypted_attr"`
	ImageEncryptedAttr               []string               `json:"image_encrypted_attr" yaml:"image_encrypted_attr"`
	ImageRestrictedAttr              []string               `json:"image_restricted_attr" yaml:"image_restricted_attr"`
	ImMAD                            []config.ImMAD         `json:"im_mad" yaml:"im_mad"`
	InheritDatastoreAttr             []string               `json:"inherit_datastore_attr" yaml:"inherit_datastore_attr"`
	InheritImageAttr                 []string               `json:"inherit_image_attr" yaml:"inherit_image_attr"`
	InheritVnetAttr                  []string               `json:"inherit_vnet_attr" yaml:"inherit_vnet_attr"`
	IpamMAD                          []config.IpamMAD       `json:"ipam_mad" yaml:"ipam_mad"`
	KeepaliveMaxConn                 []int                  `json:"keepalive_max_conn" yaml:"keepalive_max_conn"`
	KeepaliveTimeout                 []int                  `json:"keepalive_timeout" yaml:"keepalive_timeout"`
	ListenAddress                    []string               `json:"listen_address" yaml:"listen_address"`
	Log                              []config.Log           `json:"log" yaml:"log"`
	LogCallFormat                    []string               `json:"log_call_format" yaml:"log_call_format"`
	MACPrefix                        []string               `json:"macprefix" yaml:"macprefix"`
	ManagerTimer                     []int                  `json:"manager_timer" yaml:"manager_timer"`
	MarketMAD                        []config.MarketMAD     `json:"market_mad" yaml:"market_mad"`
	MarketMADConf                    []config.MarketMADConf `json:"market_madconf" yaml:"market_madconf"`
	MaxConn                          int                    `json:"max_conn" yaml:"max_conn"`
	MaxConnBacklog                   int                    `json:"max_conn_backlog" yaml:"max_conn_backlog"`
	MessageSize                      int                    `json:"message_size" yaml:"message_size"`
	MonitoringIntervalDatastore      int                    `json:"monitoring_interval_datastore" yaml:"monitoring_interval_datastore"`
	MonitoringIntervalDBUpdate       int                    `json:"monitoring_interval_dbupdate" yaml:"monitoring_interval_dbupdate"`
	MonitoringIntervalHost           int                    `json:"monitoring_interval_host" yaml:"monitoring_interval_host"`
	MonitoringIntervalMarket         int                    `json:"monitoring_interval_market" yaml:"monitoring_interval_market"`
	MonitoringIntervalInstance       int                    `json:"monitoring_interval_instance" yaml:"monitoring_interval_instance"`
	NetworkSize                      int                    `json:"network_size" yaml:"network_size"`
	Key                              []string               `json:"one_key" yaml:"one_key"`
	PCIPassthroughBus                string                 `json:"pcipassthrough_bus" yaml:"pcipassthrough_bus"`
	Port                             int                    `json:"port" yaml:"port"`
	Raft                             config.Raft            `json:"raft" yaml:"raft"`
	RPCLog                           string                 `json:"rpclog" yaml:"rpclog"`
	ScriptsRemoteDir                 string                 `json:"scripts_remote_dir" yaml:"scripts_remote_dir"`
	SessionExpirationTime            int                    `json:"session_expiration_time" yaml:"session_expiration_time"`
	ShowbackOnlyRunning              string                 `json:"showback_only_running" yaml:"showback_only_running"`
	Timeout                          int                    `json:"timeout" yaml:"timeout"`
	TmMAD                            []config.TmMAD         `json:"tm_mad" yaml:"tm_mad"`
	TmMADConf                        []config.TmMADConf     `json:"tm_madconf" yaml:"tm_madconf"`
	UserEncryptedAttr                []string               `json:"user_encrypted_attr" yaml:"user_encrypted_attr"`
	UserRestrictedAttr               []string               `json:"user_restricted_attr" yaml:"user_restricted_attr"`
	VLANIDs                          config.VLANIDs         `json:"vlanids" yaml:"vlanids"`
	InstanceAdminOperations          string                 `json:"instance_admin_operations" yaml:"instance_admin_operations"`
	InstanceEncryptedAttr            []string               `json:"instance_encrypted_attr" yaml:"instance_encrypted_attr"`
	InstanceMAD                      []config.InstanceMAD   `json:"instance_mad" yaml:"instance_mad"`
	InstanceManageOperations         string                 `json:"instance_manage_operations" yaml:"instance_manage_operations"`
	InstanceMonitoringExpirationTime []string               `json:"instance_monitoring_expiration_time" yaml:"instance_monitoring_expiration_time"`
	InstanceRestrictedAttr           []string               `json:"instance_restricted_attr" yaml:"instance_restricted_attr"`
	InstanceSnapshotFactor           string                 `json:"instance_snapshot_factor" yaml:"instance_snapshot_factor"`
	InstanceSubmitOnHold             string                 `json:"instance_submit_on_hold" yaml:"instance_submit_on_hold"`
	InstanceUseOperations            []string               `json:"instance_use_operations" yaml:"instance_use_operations"`
	VNCPorts                         config.VNCPorts        `json:"vncports" yaml:"vncports"`
	VNetEncryptedAttr                []string               `json:"vnet_encrypted_attr" yaml:"vnet_encrypted_attr"`
	VNetRestrictedAttr               []string               `json:"vnet_restricted_attr" yaml:"vnet_restricted_attr"`
	VNMADConf                        []config.VNetMADConf   `json:"vnmadconf" yaml:"vnmadconf"`
	VxlanIDs                         config.VxlanIDs        `json:"vxlan_ids" yaml:"vxlan_ids"`
}

// Hostname returns the hostname of the instance.
func (i *Instance) Hostname() string {
	l := len(i.HistoryRecords)
	if l == 0 {
		return ""
	}

	return i.HistoryRecords[l-1].Hostname
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (i *Instance) ParseTemplate() (*instance.Template, error) {
	return instance.ParseTemplate(i.Template)
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (c *Cluster) ParseTemplate() (*ClusterTemplate, error) {
	var t ClusterTemplate

	for key, value := range c.Template {
		if v, ok := value.(string); ok {
			switch key {
			case "RESERVED_CPU":
				t.ReservedCPU = v
			case "RESERVED_MEM":
				t.ReservedMemory = v
			}
		}
	}

	return &t, nil
}
