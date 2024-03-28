// Package instance contains structs for the Instance payload.
package instance

import (
	"time"

	"github.com/softiron/manifold-api/deprecated/v1/hc"
	"github.com/softiron/manifold-api/internal/api"
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

// V1 returns the v1 hc.InstanceState that corresponds to the State.
func (s State) V1() hc.InstanceState {
	return hc.InstanceState(s)
}

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
	AutomaticDSRequirements   string
	AutomaticNICRequirements  string
	AutomaticRequirements     string
	CPU                       float32
	CPUCost                   string
	CPUModel                  CPUModel
	CloningTemplateID         string
	Context                   Context
	CreatedBy                 string
	Disk                      []Disk
	DiskCost                  string
	Emulator                  string
	Features                  string
	Graphics                  Graphics
	HypervOptions             string
	Imported                  string
	Input                     string
	InstanceGroup             []string `template:"VMGROUP"`
	InstanceID                int      `template:"VMID"`
	Memory                    int
	MemoryCost                string
	MemoryMax                 string
	MemoryResizeMode          string
	MemorySlots               string
	NIC                       []NIC
	NICAlias                  []NICAlias
	NICDefault                string
	NumaNode                  string
	OS                        OS
	PCI                       string
	Raw                       string
	SchedAction               []SchedAction
	SecurityGroupRule         []string
	Snapshot                  []Snapshot
	SpiceOptions              string
	SubmitOnHold              string
	TemplateID                int
	TmMADSystem               string
	Topology                  string
	VCPU                      int
	VCPUMax                   string
	VRouterID                 string
	VRouterKeepalivedID       string
	VRouterKeepalivedPassword string
}

// CPUModel is the API payload based on the legacy xmlrpc backend.
type CPUModel struct {
	Model string
}

// Context is the API payload based on the legacy xmlrpc backend.
type Context struct {
	DiskID       int
	Firmware     string
	GuestOS      string
	Network      bool
	SSHPublicKey string
	Target       string
}

// OS is the API payload based on the legacy xmlrpc backend.
type OS struct {
	UUID    string
	Arch    string
	Boot    string
	DiskBus string
}

// History is the API payload based on the legacy xmlrpc backend.
type History struct {
	ID                 int      `json:"id" yaml:"id"`
	SequenceNumber     int      `json:"sequence_number" yaml:"sequence_number"`
	Hostname           string   `json:"hostname" yaml:"hostname"`
	HostID             int      `json:"host_id" yaml:"host_id"`
	ClusterID          int      `json:"cluster_id" yaml:"cluster_id"`
	StartTime          api.Time `json:"start_time" yaml:"start_time"`
	EndTime            api.Time `json:"end_time" yaml:"end_time"`
	InstanceMAD        string   `json:"instance_mad" yaml:"instance_mad"`
	TransferManagerMAD string   `json:"transfer_manager_mad" yaml:"transfer_manager_mad"`
	DatastoreID        int      `json:"datastore_id" yaml:"datastore_id"`
	PrologStartTime    api.Time `json:"prolog_start_time" yaml:"prolog_start_time"`
	PrologEndtime      api.Time `json:"prolog_end_time" yaml:"prolog_end_time"`
	RunningStartTime   api.Time `json:"running_start_time" yaml:"running_start_time"`
	RunningEndTime     api.Time `json:"running_end_time" yaml:"running_end_time"`
	EpilogStartTime    api.Time `json:"epilog_start_time" yaml:"epilog_start_time"`
	EpilogEndTime      api.Time `json:"epilog_end_time" yaml:"epilog_end_time"`
	Action             Action   `json:"action" yaml:"action"`
	UserID             int      `json:"user_id" yaml:"user_id"`
	GroupID            int      `json:"group_id" yaml:"group_id"`
	RequestID          int      `json:"request_id" yaml:"request_id"`
}

// Monitoring is the API payload based on the legacy xmlrpc backend.
type Monitoring struct {
	CPU                             float64    `json:"cpu" yaml:"cpu"`
	DiskReadBytes                   int        `json:"disk_read_bytes" yaml:"disk_read_bytes"`
	DiskReadIOps                    int        `json:"disk_read_iops" yaml:"disk_read_iops"`
	DiskWriteBytes                  int        `json:"disk_write_bytes" yaml:"disk_write_bytes"`
	DiskWriteIOps                   int        `json:"disk_write_iops" yaml:"disk_write_iops"`
	DiskSize                        []DiskSize `json:"disk_size" yaml:"disk_size"`
	ID                              int        `json:"id" yaml:"id"`
	Memory                          int        `json:"memory" yaml:"memory"`
	NetRX                           int        `json:"netrx" yaml:"netrx"`
	NetTX                           int        `json:"nettx" yaml:"nettx"`
	Timestamp                       api.Time   `json:"timestamp" yaml:"timestamp"`
	VCenterESXHost                  string     `json:"vcenter_esxhost" yaml:"vcenter_esxhost"`
	VCenterGuestState               string     `json:"vcenter_guest_state" yaml:"vcenter_guest_state"`
	VCenterRPName                   string     `json:"vcenter_rpname" yaml:"vcenter_rpname"`
	VCenterVMWareToolsRunningStatus string     `json:"vcenter_vmware_tools_running_status" yaml:"vcenter_vmware_tools_running_status"`
	VCenterVMWareToolsVersion       string     `json:"vcenter_vmware_tools_version" yaml:"vcenter_vmware_tools_version"`
	VCenterVMWareToolsVersionStatus string     `json:"vcenter_vmware_tools_version_status" yaml:"vcenter_vmware_tools_version_status"`
	VCenterVMName                   string     `json:"vcenter_vmname" yaml:"vcenter_vmname"`
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
	Snapshot     []ImageSnapshot `json:"snapshot" yaml:"snapshot"`
}

// DiskSnapshots is a Snapshots for a specific disk.
type DiskSnapshots struct {
	Snapshots
	DiskID int `json:"disk_id" yaml:"disk_id"`
}

// Disk is the API payload based on the legacy xmlrpc backend.
type Disk struct {
	Format                string
	AllowOrphans          string
	Clone                 bool
	CloneTarget           string
	ClusterID             int
	Datastore             string
	DatastoreID           int
	DevPrefix             string
	DiskID                int
	DiskSnapshotTotalSize int
	DiskType              string
	Driver                string
	Image                 string
	ImageID               int
	ImageState            int
	ImageUserName         string
	LnTarget              string
	OriginalSize          int
	Persistent            bool
	PoolName              string
	Readonly              bool
	Save                  bool
	Size                  int
	Source                string
	Target                string
	TmMAD                 string
	TmMADSystem           string
	Type                  string
}

// Graphics is the API payload based on the legacy xmlrpc backend.
type Graphics struct {
	Listen   string
	Port     int
	Type     string
	Password string
	Keymap   string
}

// NIC is the API payload based on the legacy xmlrpc backend.
type NIC struct {
	ID             int
	NetworkID      int
	IP             string
	MAC            string
	Model          string
	VirtioQueues   string
	Phydev         string
	SecurityGroups string
}

// NICAlias is the API payload based on the legacy xmlrpc backend.
type NICAlias struct {
	AliasID  string
	Parent   string
	ParentID string
}

// SchedAction is the API payload based on the legacy xmlrpc backend.
type SchedAction struct {
	Action   string
	Args     string
	Days     string
	EndType  string
	EndValue string
	ID       string
	Repeat   string
	Time     string
	Warning  string
}

// Snapshot is the API payload based on the legacy xmlrpc backend.
type Snapshot struct {
	Action         string
	Active         bool
	HypervisorID   string
	Name           string
	SnapshotID     int
	SystemDiskSize int
	Time           time.Time
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
