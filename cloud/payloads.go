package cloud

// Top level payloads from the HyperCloud backend API.

import (
	"fmt"
	"strconv"

	"github.com/softiron/hypercloud-api/cloud/config"
	"github.com/softiron/hypercloud-api/cloud/instance"
	"github.com/softiron/hypercloud-api/cloud/insttmpl"
	"github.com/softiron/hypercloud-api/cloud/nettmpl"
	"github.com/softiron/hypercloud-api/internal/time"
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
	ID           int             `json:"id" yaml:"id"`
	Name         string          `json:"name" yaml:"name"`
	Hosts        []int           `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Datastores   []int           `json:"datastores,omitempty" yaml:"datastores,omitempty"`
	Networks     []int           `json:"networks,omitempty" yaml:"networks,omitempty"`
	Template     ClusterTemplate `json:"template" yaml:"template"`
	TemplateText string          `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// ClusterTemplate is the API payload based on the legacy xmlrpc backend.
type ClusterTemplate struct {
	Values         map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	ReservedCPU    string            `json:"reserved_cpu,omitempty" yaml:"reserved_cpu,omitempty"`
	ReservedMemory string            `json:"reserved_mem,omitempty" yaml:"reserved_mem,omitempty"`
}

// Document is the API payload based on the legacy xmlrpc backend.
type Document struct {
	ID           int              `json:"id" yaml:"id"`
	UserID       int              `json:"user_id" yaml:"user_id"`
	GroupID      int              `json:"group_id" yaml:"group_id"`
	UserName     string           `json:"user_name" yaml:"user_name"`
	GroupName    string           `json:"group_name" yaml:"group_name"`
	Name         string           `json:"name" yaml:"name"`
	Type         string           `json:"type" yaml:"type"`
	Permissions  Permissions      `json:"permissions" yaml:"permissions"`
	Lock         Lock             `json:"lock" yaml:"lock"`
	Template     DocumentTemplate `json:"template" yaml:"template"`
	TemplateText string           `json:"template_text" yaml:"template_text"`
}

// DocumentTemplate is the API payload based on the legacy xmlrpc backend.
type DocumentTemplate struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}

// Instance is the API payload based on the legacy xmlrpc backend.
type Instance struct {
	ID               int                      `json:"id" yaml:"id"`
	UserID           int                      `json:"user_id" yaml:"user_id"`
	GroupID          int                      `json:"group_id" yaml:"group_id"`
	UserName         string                   `json:"user_name" yaml:"user_name"`
	GroupName        string                   `json:"group_name" yaml:"group_name"`
	Name             string                   `json:"name" yaml:"name"`
	Permissions      Permissions              `json:"permissions" yaml:"permissions"`
	LastPoll         time.Time                `json:"last_poll,omitempty" yaml:"last_poll,omitempty"`
	State            instance.State           `json:"state" yaml:"state"`
	LCMState         LCMState                 `json:"lcm_state" yaml:"lcm_state"`
	PrevState        instance.State           `json:"prev_state" yaml:"prev_state"`
	PrevLCMState     LCMState                 `json:"prev_lcm_state" yaml:"prev_lcm_state"`
	Reschedule       bool                     `json:"reschedule,omitempty" yaml:"reschedule,omitempty"`
	StartTime        time.Time                `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	EndTime          time.Time                `json:"end_time,omitempty" yaml:"end_time,omitempty"`
	DeployID         string                   `json:"deploy_id" yaml:"deploy_id"`
	Monitoring       instance.Monitoring      `json:"monitoring" yaml:"monitoring"`
	RawTemplate      Template                 `json:"raw_template,omitempty" yaml:"raw_template,omitempty"`
	Template         instance.Template        `json:"template" yaml:"template"`
	TemplateText     string                   `json:"template_text" yaml:"template_text"`
	RawUserTemplate  map[string]any           `json:"raw_user_template,omitempty" yaml:"raw_user_template,omitempty"`
	UserTemplate     instance.UserTemplate    `json:"user_template" yaml:"user_template"`
	UserTemplateText string                   `json:"user_template_text" yaml:"user_template_text"`
	HistoryRecords   []instance.History       `json:"history_records,omitempty" yaml:"history_records,omitempty"`
	Snapshots        []instance.DiskSnapshots `json:"snapshots,omitempty" yaml:"snapshots,omitempty"`
}

// Hostname returns the hostname of the instance.
func (i *Instance) Hostname() string {
	l := len(i.HistoryRecords)
	if l == 0 {
		return ""
	}

	return i.HistoryRecords[l-1].Hostname
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
	TemplateText    string      `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// MarketplaceApp is the API payload based on the legacy xmlrpc backend.
type MarketplaceApp struct {
	ID            int                    `json:"id" yaml:"id"`
	UID           int                    `json:"uid" yaml:"uid"`
	GID           int                    `json:"gid" yaml:"gid"`
	UserName      string                 `json:"user_name" yaml:"user_name"`
	GroupName     string                 `json:"group_name" yaml:"group_name"`
	Lock          Lock                   `json:"lock" yaml:"lock"`
	Regtime       int                    `json:"regtime" yaml:"regtime"`
	Name          string                 `json:"name" yaml:"name"`
	ZoneID        string                 `json:"zone_id" yaml:"zone_id"`
	OriginID      string                 `json:"origin_id" yaml:"origin_id"`
	Source        string                 `json:"source" yaml:"source"`
	MD5           string                 `json:"md5" yaml:"md5"`
	Size          int                    `json:"size" yaml:"size"`
	Description   string                 `json:"description" yaml:"description"`
	Version       string                 `json:"version" yaml:"version"`
	Format        string                 `json:"format" yaml:"format"`
	AppTemplate64 string                 `json:"apptemplate64" yaml:"apptemplate64"`
	MarketplaceID int                    `json:"marketplace_id" yaml:"marketplace_id"`
	Marketplace   string                 `json:"marketplace" yaml:"marketplace"`
	State         int                    `json:"state" yaml:"state"`
	Type          int                    `json:"type" yaml:"type"`
	Permissions   Permissions            `json:"permissions" yaml:"permissions"`
	Template      MarketplaceAppTemplate `json:"template" yaml:"template"`
	TemplateText  string                 `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// MarketplaceAppTemplate is the API payload based on the legacy xmlrpc backend.
type MarketplaceAppTemplate struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
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
	ID           int                   `json:"id" yaml:"id"`
	UID          int                   `json:"uid" yaml:"uid"`
	GID          int                   `json:"gid" yaml:"gid"`
	UserName     string                `json:"user_name" yaml:"user_name"`
	GroupName    string                `json:"group_name" yaml:"group_name"`
	Name         string                `json:"name" yaml:"name"`
	Permissions  Permissions           `json:"permissions" yaml:"permissions"`
	Lock         Lock                  `json:"lock" yaml:"lock"`
	Roles        []instance.GroupRole  `json:"roles" yaml:"roles"`
	Template     InstanceGroupTemplate `json:"template" yaml:"template"`
	TemplateText string                `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// InstanceGroupTemplate is the API payload based on the legacy xmlrpc backend.
type InstanceGroupTemplate struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}

// InstanceTemplate is the API payload based on the legacy xmlrpc backend.
type InstanceTemplate struct {
	ID           int               `json:"id" yaml:"id"`
	UID          int               `json:"uid" yaml:"uid"`
	GID          int               `json:"gid" yaml:"gid"`
	Uname        string            `json:"uname" yaml:"uname"`
	Gname        string            `json:"gname" yaml:"gname"`
	Name         string            `json:"name" yaml:"name"`
	Lock         Lock              `json:"lock" yaml:"lock"`
	Permissions  Permissions       `json:"permissions" yaml:"permissions"`
	Regtime      int               `json:"regtime" yaml:"regtime"`
	Template     insttmpl.Template `json:"template" yaml:"template"`
	TemplateText string            `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// NetworkTemplate is the API payload based on the legacy xmlrpc backend.
type NetworkTemplate struct {
	ID           int              `json:"id" yaml:"id"`
	UID          int              `json:"uid" yaml:"uid"`
	GID          int              `json:"gid" yaml:"gid"`
	UserName     string           `json:"user_name" yaml:"user_name"`
	GroupName    string           `json:"group_name" yaml:"group_name"`
	Name         string           `json:"name" yaml:"name"`
	Lock         Lock             `json:"lock" yaml:"lock"`
	Permissions  Permissions      `json:"permissions" yaml:"permissions"`
	Regtime      int              `json:"regtime" yaml:"regtime"`
	Template     nettmpl.Template `json:"template" yaml:"template"`
	TemplateText string           `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// Router is the API payload based on the legacy xmlrpc backend.
type Router struct {
	ID           int            `json:"id" yaml:"id"`
	UID          int            `json:"uid" yaml:"uid"`
	GID          int            `json:"gid" yaml:"gid"`
	UserName     string         `json:"user_name" yaml:"user_name"`
	GroupName    string         `json:"group_name" yaml:"group_name"`
	Name         string         `json:"name" yaml:"name"`
	Permissions  Permissions    `json:"permissions" yaml:"permissions"`
	Lock         Lock           `json:"lock" yaml:"lock"`
	Instances    []int          `json:"instances" yaml:"instances"`
	Template     RouterTemplate `json:"template" yaml:"template"`
	TemplateText string         `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// RouterTemplate is the API payload based on the legacy xmlrpc backend.
type RouterTemplate struct {
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}

// HyperCloudConfiguration is the API payload based on the legacy xmlrpc backend.
type HyperCloudConfiguration struct {
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

// Template returns a structured subset of the nested key x value pair map.
func (t Template) Template() (*instance.Template, error) { // nolint:gocognit
	var dst instance.Template

	for key, value := range t {
		switch v := value.(type) {
		case string:
			switch key {
			case "AUTOMATIC_DS_REQUIREMENTS":
				dst.AutomaticDSRequirements = v
			case "AUTOMATIC_NIC_REQUIREMENTS":
				dst.AutomaticNICRequirements = v
			case "AUTOMATIC_REQUIREMENTS":
				dst.AutomaticRequirements = v
			case "CPU":
				n, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return nil, fmt.Errorf("invalid CPU value %q: %w", v, err)
				}
				dst.CPU = float32(n)
			case "CPU_COST":
				dst.CPUCost = v
			case "CLONING_TEMPLATE_ID":
				dst.CloningTemplateID = v
			case "CREATED_BY":
				dst.CreatedBy = v
			case "DISK_COST":
				dst.DiskCost = v
			case "EMULATOR":
				dst.Emulator = v
			case "FEATURES":
				dst.Features = v
			case "HYPERV_OPTIONS":
				dst.HypervOptions = v
			case "IMPORTED":
				dst.Imported = v
			case "INPUT":
				dst.Input = v
			case "MEMORY":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid MEMORY value %q: %w", v, err)
				}
				dst.Memory = n
			case "MEMORY_COST":
				dst.MemoryCost = v
			case "MEMORY_MAX":
				dst.MemoryMax = v
			case "NIC_DEFAULT":
				dst.NICDefault = v
			case "NUMA_NODE":
				dst.NumaNode = v
			case "PCI":
				dst.PCI = v
			case "RAW":
				dst.Raw = v
			case "SPICE_OPTIONS":
				dst.SpiceOptions = v
			case "SUBMIT_ON_HOLD":
				dst.SubmitOnHold = v
			case "TEMPLATE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid TEMPLATE_ID value %q: %w", v, err)
				}
				dst.TemplateID = n
			case "TM_MAD_SYSTEM":
				dst.TmMADSystem = v
			case "TOPOLOGY":
				dst.Topology = v
			case "VCPU":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid VCPU value %q: %w", v, err)
				}
				dst.VCPU = n
			case "VCPU_MAX":
				dst.VCPUMax = v
			case "VMID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid VMID value %q: %w", v, err)
				}
				dst.InstanceID = n
			case "VROUTER_ID":
				dst.VRouterID = v
			case "VROUTER_KEEPALIVED_ID":
				dst.VRouterKeepalivedID = v
			case "VROUTER_KEEPALIVED_PASSWORD":
				dst.VRouterKeepalivedPassword = v
			}
		case map[string]any:
			switch key {
			case "CONTEXT":
				c, err := newInstanceContext(v)
				if err != nil {
					return nil, err
				}
				dst.Context = *c
			case "CPU_MODEL":
				m := newInstanceCPUModel(v)
				dst.CPUModel = *m
			case "DISK":
				d, err := newInstanceDisk(v)
				if err != nil {
					return nil, err
				}
				dst.Disk = []instance.Disk{*d}
			case "GRAPHICS":
				g, err := newInstanceGraphics(v)
				if err != nil {
					return nil, err
				}
				dst.Graphics = *g
			case "NIC":
				n, err := newInstanceNIC(v)
				if err != nil {
					return nil, err
				}
				dst.NIC = []instance.NIC{*n}
			case "NIC_ALIAS":
				n := newInstanceNICAlias(v)
				dst.NICAlias = []instance.NICAlias{*n}
			case "OS":
				o := newInstanceOS(v)
				dst.OS = *o
			case "SCHED_ACTION":
				s := newInstanceSchedAction(v)
				dst.SchedAction = []instance.SchedAction{*s}
			case "SNAPSHOT":
				s, err := newInstanceSnapshot(v)
				if err != nil {
					return nil, err
				}
				dst.Snapshot = []instance.Snapshot{*s}
			}
		case []map[string]any:
			switch key {
			case "DISK":
				for _, m := range v {
					d, err := newInstanceDisk(m)
					if err != nil {
						return nil, err
					}
					dst.Disk = append(dst.Disk, *d)
				}
			case "NIC":
				for _, m := range v {
					n, err := newInstanceNIC(m)
					if err != nil {
						return nil, err
					}
					dst.NIC = append(dst.NIC, *n)
				}
			case "NIC_ALIAS":
				for _, m := range v {
					n := newInstanceNICAlias(m)
					dst.NICAlias = append(dst.NICAlias, *n)
				}
			case "SCHED_ACTION":
				for _, m := range v {
					s := newInstanceSchedAction(m)
					dst.SchedAction = append(dst.SchedAction, *s)
				}
			case "SNAPSHOT":
				for _, m := range v {
					s, err := newInstanceSnapshot(m)
					if err != nil {
						return nil, err
					}
					dst.Snapshot = append(dst.Snapshot, *s)
				}
			}
		}
	}

	return &dst, nil
}

func newInstanceContext(m map[string]any) (*instance.Context, error) {
	var dst instance.Context

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "DISK_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DISK_ID value %q: %w", v, err)
				}
				dst.DiskID = n
			case "FIRMWARE":
				dst.Firmware = v
			case "GUESTOS":
				dst.GuestOS = v
			case "NETWORK":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid NETWORK value %q: %w", v, err)
				}
				dst.Network = b
			case "SSH_PUBLIC_KEY":
				dst.SSHPublicKey = v
			case "TARGET":
				dst.Target = v
			}
		}
	}

	return &dst, nil
}

func newInstanceCPUModel(m map[string]any) *instance.CPUModel {
	var dst instance.CPUModel

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key { // nolint:gocritic
			case "MODEL":
				dst.Model = v
			}
		}
	}

	return &dst
}

func newInstanceGraphics(m map[string]any) (*instance.Graphics, error) {
	var dst instance.Graphics

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "KEYMAP":
				dst.Keymap = v
			case "LISTEN":
				dst.Listen = v
			case "PORT":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid PORT value %q: %w", v, err)
				}
				dst.Port = n
			case "TYPE":
				dst.GraphicsType = v
			case "PASSWD":
				dst.Password = v
			}
		}
	}

	return &dst, nil
}

func newInstanceOS(m map[string]any) *instance.OS {
	var dst instance.OS

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ARCH":
				dst.Arch = v
			case "BOOT":
				dst.Boot = v
			case "UUID":
				dst.UUID = v
			case "SD_DISK_BUS":
				dst.DiskBus = v
			}
		}
	}

	return &dst
}

func newInstanceDisk(m map[string]any) (*instance.Disk, error) { // nolint:gocognit
	var dst instance.Disk

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ALLOW_ORPHANS":
				dst.AllowOrphans = v
			case "CLONE":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid CLONE value %q: %w", v, err)
				}
				dst.Clone = b
			case "CLONE_TARGET":
				dst.CloneTarget = v
			case "CLUSTER_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid CLUSTER_ID value %q: %w", v, err)
				}
				dst.ClusterID = n
			case "DATASTORE":
				dst.Datastore = v
			case "DATASTORE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DATASTORE_ID value %q: %w", v, err)
				}
				dst.DatastoreID = n
			case "DEV_PREFIX":
				dst.DevPrefix = v
			case "DISK_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DISK_ID value %q: %w", v, err)
				}
				dst.DiskID = n
			case "DISK_SNAPSHOT_TOTAL_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DISK_SNAPSHOT_TOTAL_SIZE value %q: %w", v, err)
				}
				dst.DiskSnapshotTotalSize = n
			case "DISK_TYPE":
				dst.DiskType = v
			case "DRIVER":
				dst.Driver = v
			case "FORMAT":
				dst.Format = v
			case "IMAGE":
				dst.Image = v
			case "IMAGE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid IMAGE_ID value %q: %w", v, err)
				}
				dst.ImageID = n
			case "IMAGE_STATE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid IMAGE_STATE value %q: %w", v, err)
				}
				dst.ImageState = n
			case "IMAGE_UNAME":
				dst.ImageUserName = v
			case "LN_TARGET":
				dst.LnTarget = v
			case "ORIGINAL_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid ORIGINAL_SIZE value %q: %w", v, err)
				}
				dst.OriginalSize = n
			case "PERSISTENT":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid PERSISTENT value %q: %w", v, err)
				}
				dst.Persistent = b
			case "POOL_NAME":
				dst.PoolName = v
			case "READONLY":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid READONLY value %q: %w", v, err)
				}
				dst.Readonly = b
			case "SAVE":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SAVE value %q: %w", v, err)
				}
				dst.Save = b
			case "SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SIZE value %q: %w", v, err)
				}
				dst.Size = n
			case "SOURCE":
				dst.Source = v
			case "TARGET":
				dst.Target = v
			case "TM_MAD":
				dst.TmMAD = v
			case "TM_MAD_SYSTEM":
				dst.TmMADSystem = v
			case "TYPE":
				dst.Type = v
			case "VCENTER_DS_REF":
				dst.VCenterDSRef = v
			case "VCENTER_INSTANCE_ID": // nolint:goconst
				dst.VCenterInstanceID = v
			}
		}
	}

	return &dst, nil
}

func newInstanceNIC(m map[string]any) (*instance.NIC, error) {
	var dst instance.NIC

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			// "AR_ID": "0",
			// "BRIDGE": "mgmt0",
			// "BRIDGE_TYPE": "linux",
			// "CLUSTER_ID": "0",
			// "GATEWAY": "192.168.8.1",
			// "NAME": "NIC0",
			// "NETWORK": "Infrastructure Management Network",
			// "TARGET": "one-9-0",
			// "VLAN_ID": "12",
			// "VN_MAD": "802.1Q"
			case "NIC_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid NIC_ID value %q: %w", v, err)
				}
				dst.ID = n
			case "NETWORK_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid NETWORK_ID value %q: %w", v, err)
				}
				dst.NetworkID = n
			case "IP":
				dst.IP = v
			case "MAC":
				dst.MAC = v
			case "MODEL":
				dst.Model = v
			case "VIRTIO_QUEUES":
				dst.VirtioQueues = v
			case "PhyDev":
				dst.Phydev = v
			case "SECURITY_GROUPS":
				dst.SecurityGroups = v
			case "VCENTER_INSTANCE_ID":
				dst.VCenterInstanceID = v
			case "VCENTER_NET_REF":
				dst.VCenterNetRef = v
			case "VCENTER_PORTGROUP_TYPE":
				dst.VCenterPortgroupType = v
			}
		}
	}

	return &dst, nil
}

func newInstanceNICAlias(m map[string]any) *instance.NICAlias {
	var dst instance.NICAlias

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ALIAS_ID":
				dst.AliasID = v
			case "PARENT":
				dst.Parent = v
			case "PARENT_ID":
				dst.ParentID = v
			case "VCENTER_INSTANCE_ID":
				dst.VCenterInstanceID = v
			case "VCENTER_NET_REF":
				dst.VCenterNetRef = v
			case "VCENTER_PORTGROUP_TYPE":
				dst.VCenterPortgroupType = v
			}
		}
	}

	return &dst
}

func newInstanceSchedAction(m map[string]any) *instance.SchedAction {
	var dst instance.SchedAction

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ACTION":
				dst.Action = v
			case "ARGS":
				dst.Args = v
			case "DAYS":
				dst.Days = v
			case "END_TYPE":
				dst.EndType = v
			case "END_VALUE":
				dst.EndValue = v
			case "ID":
				dst.ID = v
			case "REPEAT":
				dst.Repeat = v
			case "TIME":
				dst.Time = v
			case "WARNING":
				dst.Warning = v
			}
		}
	}

	return &dst
}

func newInstanceSnapshot(m map[string]any) (*instance.Snapshot, error) {
	var dst instance.Snapshot

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ACTION":
				dst.Action = v
			case "ACTIVE":
				b, err := str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid ACTIVE value %q: %w", v, err)
				}
				dst.Active = b
			case "HYPERVISOR_ID":
				dst.HypervisorID = v
			case "NAME":
				dst.Name = v
			case "SNAPSHOT_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SNAPSHOT_ID value %q: %w", v, err)
				}
				dst.SnapshotID = n
			case "SYSTEM_DISK_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SYSTEM_DISK_SIZE value %q: %w", v, err)
				}
				dst.SystemDiskSize = n
			case "TIME":
				t, err := strconv.ParseInt(v, 10, 0)
				if err != nil {
					return nil, fmt.Errorf("invalid TIME value %q: %w", v, err)
				}
				dst.Time = time.Unix(t, 0) // TODO: we don't print as json so this should be time.Time (not our wrapper type)
			}
		}
	}

	return &dst, nil
}
