// Package instance contains structs for the Instance payload.
package instance

import (
	"encoding/xml"
	"time"

	"github.com/softiron/hypercloud-api/cloud/template"
)

//go:generate go run "github.com/dmarkham/enumer" -type State -linecomment -text

// State represents the state of an instance.
type State int

// State values.
const (
	InitState      State = iota // init
	PendingState                // pending
	HoldState                   // hold
	ActiveState                 // active
	StoppedState                // stopped
	SuspendedState              // suspended
	DoneState                   // done
	_
	OffState                      // off
	UndeployedState               // undeployed
	CloningState                  // cloning
	CloningFailedState            // cloning_failed
	AnyState           State = -2 // all
	NotDoneState       State = -1 // not_done
)

//go:generate go run "github.com/dmarkham/enumer" -type Action -linecomment -text

// Action is the action to perform on an intstance.
type Action int

// Action values.
const (
	NoneAction               Action = iota // none
	MigrateAction                          // migrate
	LiveMigrateAction                      // live-migrate
	ShutdownAction                         // shutdown
	ShutdownHardAction                     // shutdown-hard
	UndeployAction                         // undeploy
	UndeployHardAction                     // undeploy-hard
	HoldAction                             // hold
	ReleaseAction                          // release
	StopAction                             // stop
	SuspendAction                          // suspend
	ResumeAction                           // resume
	BootAction                             // boot
	DeleteAction                           // delete
	DeleteRecreateAction                   // delete-recreate
	RebootAction                           // reboot
	RebootHardAction                       // reboot-hard
	RescheduleAction                       // resched
	UnrescheduleAction                     // unresched
	PowerOffAction                         // poweroff
	PowerOffHardAction                     // poweroff-hard
	DiskAttachAction                       // disk-attach
	DiskDetachAction                       // disk-detach
	NicAttachAction                        // nic-attach
	NicDetachAction                        // nic-detach
	DiskSnapshotCreateAction               // disk-snapshot-create
	DiskSnapshotDeleteAction               // disk-snapshot-delete
	TerminateAction                        // terminate
	TerminateHardAction                    // terminate-hard
	DiskResizeAction                       // disk-resize
	DeployAction                           // deploy
	ChownAction                            // chown
	ChmodAction                            // chmod
	UpdateconfAction                       // updateconf
	RenameAction                           // rename
	ResizeAction                           // resize
	UpdateAction                           // update
	SnapshotCreateAction                   // snapshot-create
	SnapshotDeleteAction                   // snapshot-delete
	SnapshotRevertAction                   // snapshot-revert
	DiskSaveasAction                       // disk-saveas
	DiskSnapshotRevertAction               // disk-snapshot-revert
	RecoverAction                          // recover
	RetryAction                            // retry
	MonitorAction                          // monitor
	DiskSnapshotRenameAction               // disk-snapshot-rename
	AliasAttachAction                      // alias-attach
	AliasDetachAction                      // alias-detach
	PoffMigrateAction                      // poff-migrate
	PoffHardMigrateAction                  // poff-hard-migrate
)

// Template is the API payload based on the legacy xmlrpc backend.
type Template struct {
	XMLName                   xml.Name        `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values                    template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	AutomaticDSRequirements   string          `json:"automatic_dsrequirements" yaml:"automatic_dsrequirements" xml:"AUTOMATIC_DSREQUIREMENTS"`
	AutomaticNICRequirements  string          `json:"automatic_nicrequirements" yaml:"automatic_nicrequirements" xml:"AUTOMATIC_NICREQUIREMENTS"`
	AutomaticRequirements     string          `json:"automatic_requirements" yaml:"automatic_requirements" xml:"AUTOMATIC_REQUIREMENTS"`
	CPU                       float32         `json:"cpu" yaml:"cpu" xml:"CPU"`
	CPUCost                   string          `json:"cpu_cost" yaml:"cpu_cost" xml:"CPU_COST"`
	CloningTemplateID         string          `json:"cloning_template_id" yaml:"cloning_template_id" xml:"CLONING_TEMPLATE_ID"`
	Context                   Context         `json:"context" yaml:"context" xml:"CONTEXT"`
	Disk                      []Disk          `json:"disk,omitempty" yaml:"disk,omitempty" xml:"DISK,omitempty"`
	DiskCost                  string          `json:"disk_cost" yaml:"disk_cost" xml:"DISK_COST"`
	Emulator                  string          `json:"emulator" yaml:"emulator" xml:"EMULATOR"`
	Features                  string          `json:"features" yaml:"features" xml:"FEATURES"`
	Graphics                  Graphics        `json:"graphics" yaml:"graphics" xml:"GRAPHICS"`
	HypervOptions             string          `json:"hyperv_options" yaml:"hyperv_options" xml:"HYPERV_OPTIONS"`
	Imported                  string          `json:"imported" yaml:"imported" xml:"IMPORTED"`
	Input                     string          `json:"input" yaml:"input" xml:"INPUT"`
	InstanceGroup             []string        `json:"instance_group,omitempty" yaml:"instance_group,omitempty" xml:"INSTANCE_GROUP,omitempty"`
	InstanceID                int             `json:"instance_id" yaml:"instance_id" xml:"INSTANCE_ID"`
	Memory                    int             `json:"memory" yaml:"memory" xml:"MEMORY"`
	MemoryCost                string          `json:"memory_cost" yaml:"memory_cost" xml:"MEMORY_COST"`
	MemoryMax                 string          `json:"memory_max" yaml:"memory_max" xml:"MEMORY_MAX"`
	MemoryResizeMode          string          `json:"memory_resize_mode" yaml:"memory_resize_mode" xml:"MEMORY_RESIZE_MODE"`
	MemorySlots               string          `json:"memory_slots" yaml:"memory_slots" xml:"MEMORY_SLOTS"`
	NIC                       []NIC           `json:"nic,omitempty" yaml:"nic,omitempty" xml:"NIC,omitempty"`
	NICAlias                  []NICAlias      `json:"nic_alias,omitempty" yaml:"nic_alias,omitempty" xml:"NIC_ALIAS,omitempty"`
	NICDefault                string          `json:"nic_default" yaml:"nic_default" xml:"NIC_DEFAULT"`
	NumaNode                  string          `json:"numa_node" yaml:"numa_node" xml:"NUMA_NODE"`
	OS                        OS              `json:"os" yaml:"os" xml:"OS"`
	PCI                       string          `json:"pci" yaml:"pci" xml:"PCI"`
	Raw                       string          `json:"raw" yaml:"raw" xml:"RAW"`
	SchedAction               []SchedAction   `json:"sched_action,omitempty" yaml:"sched_action,omitempty" xml:"SCHED_ACTION,omitempty"`
	SecurityGroupRule         []string        `json:"security_group_rule,omitempty" yaml:"security_group_rule,omitempty" xml:"SECURITY_GROUP_RULE,omitempty"`
	Snapshot                  []Snapshot      `json:"snapshot,omitempty" yaml:"snapshot,omitempty" xml:"SNAPSHOT,omitempty"`
	SpiceOptions              string          `json:"spice_options" yaml:"spice_options" xml:"SPICE_OPTIONS"`
	SubmitOnHold              string          `json:"submit_on_hold" yaml:"submit_on_hold" xml:"SUBMIT_ON_HOLD"`
	TemplateID                int             `json:"template_id" yaml:"template_id" xml:"TEMPLATE_ID"`
	TmMADSystem               string          `json:"tm_madsystem" yaml:"tm_madsystem" xml:"TM_MAD_SYSTEM"`
	Topology                  string          `json:"topology" yaml:"topology" xml:"TOPOLOGY"`
	VCPU                      int             `json:"vcpu" yaml:"vcpu" xml:"VCPU"`
	VCPUMax                   string          `json:"vcpu_max" yaml:"vcpu_max" xml:"VCPU_MAX"`
	VRouterID                 string          `json:"vrouter_id" yaml:"vrouter_id" xml:"VROUTER_ID"`
	VRouterKeepalivedID       string          `json:"vrouter_keepalived_id" yaml:"vrouter_keepalived_id" xml:"VROUTER_KEEPALIVED_ID"`
	VRouterKeepalivedPassword string          `json:"vrouter_keepalived_password" yaml:"vrouter_keepalived_password" xml:"VROUTER_KEEPALIVED_PASSWORD"`
}

// Context is the API payload based on the legacy xmlrpc backend.
type Context struct {
	Values       template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	DiskID       int             `json:"disk_id" yaml:"disk_id" xml:"DISK_ID"`
	Network      bool            `json:"network" yaml:"network" xml:"NETWORK"`
	SSHPublicKey string          `json:"ssh_public_key,omitempty" yaml:"ssh_public_key,omitempty" xml:"SSH_PUBLIC_KEY,omitempty"`
	Target       string          `json:"target" yaml:"target" xml:"TARGET"`
}

// OS is the API payload based on the legacy xmlrpc backend.
type OS struct {
	Values template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	UUID   string          `json:"uuid" yaml:"uuid" xml:"UUID"`
}

// UserTemplate is the API payload based on the legacy xmlrpc backend.
type UserTemplate struct {
	Values                template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	VCenterCCRRef         string          `json:"vcenter_ccrref,omitempty" yaml:"vcenter_ccrref,omitempty" xml:"VCENTER_CCRREF,omitempty"`
	VCenterDSRef          string          `json:"vcenter_dsref,omitempty" yaml:"vcenter_dsref,omitempty" xml:"VCENTER_DSREF,omitempty"`
	VCenterInstanceID     string          `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty" xml:"VCENTER_INSTANCE_ID,omitempty"`
	Error                 string          `json:"error" yaml:"error" xml:"ERROR"`
	Hypervisor            string          `json:"hypervisor" yaml:"hypervisor" xml:"HYPERVISOR"`
	Info                  string          `json:"info" yaml:"info" xml:"INFO"`
	Logo                  string          `json:"logo" yaml:"logo" xml:"LOGO"`
	LXDSecurityPrivileged bool            `json:"lxdsecurity_privileged" yaml:"lxdsecurity_privileged" xml:"LXDSECURITY_PRIVILEGED"`
	SchedRequirements     string          `json:"sched_requirements" yaml:"sched_requirements" xml:"SCHED_REQUIREMENTS"`
	SnapshotSchedule      string          `json:"snapshot_schedule" yaml:"snapshot_schedule" xml:"SNAPSHOT_SCHEDULE"`
}

// History is the API payload based on the legacy xmlrpc backend.
type History struct {
	ID                 int       `json:"id" yaml:"id"`
	SequenceNumber     int       `json:"sequence_number" yaml:"sequence_number"`
	Hostname           string    `json:"hostname" yaml:"hostname"`
	HostID             int       `json:"host_id" yaml:"host_id"`
	ClusterID          int       `json:"cluster_id" yaml:"cluster_id"`
	StartTime          time.Time `json:"start_time,omitempty" yaml:"start_time,omitempty"`
	EndTime            time.Time `json:"end_time,omitempty" yaml:"end_time,omitempty"`
	InstanceMAD        string    `json:"instance_mad" yaml:"instance_mad"`
	TransferManagerMAD string    `json:"transfer_manager_mad" yaml:"transfer_manager_mad"`
	DatastoreID        int       `json:"datastore_id" yaml:"datastore_id"`
	PrologStartTime    time.Time `json:"prolog_start_time,omitempty" yaml:"prolog_start_time,omitempty"`
	PrologEndtime      time.Time `json:"prolog_end_time,omitempty" yaml:"prolog_end_time,omitempty"`
	RunningStartTime   time.Time `json:"running_start_time,omitempty" yaml:"running_start_time,omitempty"`
	RunningEndTime     time.Time `json:"running_end_time,omitempty" yaml:"running_end_time,omitempty"`
	EpilogStartTime    time.Time `json:"epilog_start_time,omitempty" yaml:"epilog_start_time,omitempty"`
	EpilogEndTime      time.Time `json:"epilog_end_time,omitempty" yaml:"epilog_end_time,omitempty"`
	Action             Action    `json:"action,omitempty" yaml:"action,omitempty"`
	UserID             int       `json:"user_id" yaml:"user_id"`
	GroupID            int       `json:"group_id" yaml:"group_id"`
	RequestID          int       `json:"request_id" yaml:"request_id"`
}

// Monitoring is the API payload based on the legacy xmlrpc backend.
type Monitoring struct {
	CPU                             float64    `json:"cpu" yaml:"cpu"`
	DiskReadBytes                   int        `json:"disk_read_bytes" yaml:"disk_read_bytes"`
	DiskReadIOps                    int        `json:"disk_read_iops" yaml:"disk_read_iops"`
	DiskWriteBytes                  int        `json:"disk_write_bytes" yaml:"disk_write_bytes"`
	DiskWriteIOps                   int        `json:"disk_write_iops" yaml:"disk_write_iops"`
	DiskSize                        []DiskSize `json:"disk_size,omitempty" yaml:"disk_size,omitempty"`
	ID                              int        `json:"id" yaml:"id"`
	Memory                          int        `json:"memory" yaml:"memory"`
	NetRX                           int        `json:"netrx" yaml:"netrx"`
	NetTX                           int        `json:"nettx" yaml:"nettx"`
	Timestamp                       time.Time  `json:"timestamp" yaml:"timestamp"`
	VCenterESXHost                  string     `json:"vcenter_esxhost,omitempty" yaml:"vcenter_esxhost,omitempty"`
	VCenterGuestState               string     `json:"vcenter_guest_state,omitempty" yaml:"vcenter_guest_state,omitempty"`
	VCenterRPName                   string     `json:"vcenter_rpname,omitempty" yaml:"vcenter_rpname,omitempty"`
	VCenterVMWareToolsRunningStatus string     `json:"vcenter_vmware_tools_running_status,omitempty" yaml:"vcenter_vmware_tools_running_status,omitempty"`
	VCenterVMWareToolsVersion       string     `json:"vcenter_vmware_tools_version,omitempty" yaml:"vcenter_vmware_tools_version,omitempty"`
	VCenterVMWareToolsVersionStatus string     `json:"vcenter_vmware_tools_version_status,omitempty" yaml:"vcenter_vmware_tools_version_status,omitempty"`
	VCenterVMName                   string     `json:"vcenter_vmname,omitempty" yaml:"vcenter_vmname,omitempty"`
}

// DiskSize is the API payload based on the legacy xmlrpc backend.
type DiskSize struct {
	ID   int `json:"id" yaml:"id"`
	Size int `json:"size" yaml:"size"`
}

// Snapshots is the API payload based on the legacy xmlrpc backend.
type Snapshots struct {
	AllowOrphans string          `json:"allow_orphans" yaml:"allow_orphans"`
	CurrentBase  int             `json:"current_base" yaml:"current_base"`
	NextSnapshot int             `json:"next_snapshot" yaml:"next_snapshot"`
	Snapshot     []ImageSnapshot `json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
}

// DiskSnapshots is a Snapshots for a specific disk.
type DiskSnapshots struct {
	Snapshots
	DiskID int `json:"disk_id" yaml:"disk_id"`
}

// Disk is the API payload based on the legacy xmlrpc backend.
type Disk struct {
	Values                template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	Format                string          `json:"format" yaml:"format" xml:"FORMAT"`
	AllowOrphans          string          `json:"allow_orphans" yaml:"allow_orphans" xml:"ALLOW_ORPHANS"`
	Clone                 bool            `json:"clone" yaml:"clone" xml:"CLONE"`
	CloneTarget           string          `json:"clone_target" yaml:"clone_target" xml:"CLONE_TARGET"`
	ClusterID             int             `json:"cluster_id" yaml:"cluster_id" xml:"CLUSTER_ID"`
	Datastore             string          `json:"datastore" yaml:"datastore" xml:"DATASTORE"`
	DatastoreID           int             `json:"datastore_id" yaml:"datastore_id" xml:"DATASTORE_ID"`
	DevPrefix             string          `json:"dev_prefix" yaml:"dev_prefix" xml:"DEV_PREFIX"`
	DiskID                int             `json:"disk_id" yaml:"disk_id" xml:"DISK_ID"`
	DiskSnapshotTotalSize int             `json:"disk_snapshot_total_size" yaml:"disk_snapshot_total_size" xml:"DISK_SNAPSHOT_TOTAL_SIZE"`
	DiskType              string          `json:"disk_type" yaml:"disk_type" xml:"DISK_TYPE"`
	Driver                string          `json:"driver" yaml:"driver" xml:"DRIVER"`
	Image                 string          `json:"image" yaml:"image" xml:"IMAGE"`
	ImageID               int             `json:"image_id" yaml:"image_id" xml:"IMAGE_ID"`
	ImageState            int             `json:"image_state" yaml:"image_state" xml:"IMAGE_STATE"`
	LnTarget              string          `json:"ln_target" yaml:"ln_target" xml:"LN_TARGET"`
	OriginalSize          int             `json:"original_size" yaml:"original_size" xml:"ORIGINAL_SIZE"`
	PoolName              string          `json:"pool_name" yaml:"pool_name" xml:"POOL_NAME"`
	Readonly              bool            `json:"readonly" yaml:"readonly" xml:"READONLY"`
	Save                  bool            `json:"save" yaml:"save" xml:"SAVE"`
	Size                  int             `json:"size" yaml:"size" xml:"SIZE"`
	Source                string          `json:"source" yaml:"source" xml:"SOURCE"`
	Target                string          `json:"target" yaml:"target" xml:"TARGET"`
	TmMAD                 string          `json:"tm_mad" yaml:"tm_mad" xml:"TM_MAD"`
	TmMADSystem           string          `json:"tm_madsystem" yaml:"tm_madsystem" xml:"TM_MADSYSTEM"`
	Type                  string          `json:"type" yaml:"type" xml:"TYPE"`
	VCenterDSRef          string          `json:"vcenter_dsref,omitempty" yaml:"vcenter_dsref,omitempty" xml:"VCENTER_DS_REF,omitempty"`
	VCenterInstanceID     string          `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty" xml:"VCENTER_INSTANCE_ID,omitempty"`
}

// Graphics is the API payload based on the legacy xmlrpc backend.
type Graphics struct {
	Values       template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	Listen       string          `json:"listen" yaml:"listen" xml:"LISTEN"`
	Port         int             `json:"port" yaml:"port" xml:"PORT"`
	GraphicsType string          `json:"type" yaml:"type" xml:"TYPE"`
	Password     string          `json:"password" yaml:"password" xml:"PASSWORD"`
}

// NIC is the API payload based on the legacy xmlrpc backend.
type NIC struct {
	Values               template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	NetworkID            string          `json:"network_id" yaml:"network_id" xml:"NETWORK_ID"`
	IP                   string          `json:"ip" yaml:"ip" xml:"IP"`
	MAC                  string          `json:"mac" yaml:"mac" xml:"MAC"`
	Model                string          `json:"model" yaml:"model" xml:"MODEL"`
	VirtioQueues         string          `json:"virtio_queues" yaml:"virtio_queues" xml:"VIRTIO_QUEUES"`
	Phydev               string          `json:"phy_dev" yaml:"phy_dev" xml:"PHYDEV"`
	SecurityGroups       string          `json:"security_groups" yaml:"security_groups" xml:"SECURITY_GROUPS"`
	VCenterInstanceID    string          `json:"vcenter_instance_id" yaml:"vcenter_instance_id" xml:"VCENTER_INSTANCE_ID"`
	VCenterNetRef        string          `json:"vcenter_net_ref" yaml:"vcenter_net_ref" xml:"VCENTER_NET_REF"`
	VCenterPortgroupType string          `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type" xml:"VCENTER_PORTGROUP_TYPE"`
}

// NICAlias is the API payload based on the legacy xmlrpc backend.
type NICAlias struct {
	Values               template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	AliasID              string          `json:"alias_id" yaml:"alias_id" xml:"ALIAS_ID"`
	Parent               string          `json:"parent" yaml:"parent" xml:"PARENT"`
	ParentID             string          `json:"parent_id" yaml:"parent_id" xml:"PARENT_ID"`
	VCenterInstanceID    string          `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty" xml:"VCENTER_INSTANCE_ID,omitempty"`
	VCenterNetRef        string          `json:"vcenter_net_ref,omitempty" yaml:"vcenter_net_ref,omitempty" xml:"VCENTER_NET_REF,omitempty"`
	VCenterPortgroupType string          `json:"vcenter_portgroup_type,omitempty" yaml:"vcenter_portgroup_type,omitempty" xml:"VCENTER_PORTGROUP_TYPE,omitempty"`
}

// SchedAction is the API payload based on the legacy xmlrpc backend.
type SchedAction struct {
	Action   string `json:"action" yaml:"action" xml:"ACTION"`
	Args     string `json:"args" yaml:"args" xml:"ARGS"`
	Days     string `json:"days" yaml:"days" xml:"DAYS"`
	EndType  string `json:"end_type" yaml:"end_type" xml:"END_TYPE"`
	EndValue string `json:"end_value" yaml:"end_value" xml:"END_VALUE"`
	ID       string `json:"id" yaml:"id" xml:"ID"`
	Repeat   string `json:"repeat" yaml:"repeat" xml:"REPEAT"`
	Time     string `json:"time" yaml:"time" xml:"TIME"`
	Warning  string `json:"warning" yaml:"warning" xml:"WARNING"`
}

// Snapshot is the API payload based on the legacy xmlrpc backend.
type Snapshot struct {
	Action         string    `json:"action,omitempty" yaml:"action,omitempty" xml:"ACTION,omitempty"`
	Active         bool      `json:"active" yaml:"active" xml:"ACTIVE"`
	HypervisorID   string    `json:"hypervisor_id" yaml:"hypervisor_id" xml:"HYPERVISOR_ID"`
	Name           string    `json:"name" yaml:"name" xml:"NAME"`
	SnapshotID     int       `json:"snapshot_id" yaml:"snapshot_id" xml:"SNAPSHOT_ID"`
	SystemDiskSize int       `json:"system_disk_size" yaml:"system_disk_size" xml:"SYSTEM_DISK_SIZE"`
	Time           time.Time `json:"time,omitempty" yaml:"time,omitempty" xml:"TIME,omitempty"`
}

// ImageSnapshot is the API payload based on the legacy xmlrpc backend.
type ImageSnapshot struct {
	Active   string `json:"active" yaml:"active"`
	Children string `json:"children" yaml:"children"`
	Date     int    `json:"date" yaml:"date"`
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Parent   int    `json:"parent" yaml:"parent"`
	Size     int    `json:"size" yaml:"size"`
}

// GroupRole is the API payload based on the legacy xmlrpc backend.
type GroupRole struct {
	HostAffined     string `json:"host_affined" yaml:"host_affined"`
	HostAntiAffined string `json:"host_anti_affined" yaml:"host_anti_affined"`
	ID              int    `json:"id" yaml:"id"`
	Name            string `json:"name" yaml:"name"`
	Policy          string `json:"policy" yaml:"policy"`
}
