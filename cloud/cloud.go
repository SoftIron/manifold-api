// Package cloud documents the request and response payloads for version 2 of the sifi API.
package cloud

import "time"

// Root path for API endpoint.
const (
	AccessControlListPath = "acl"
	ClusterPath           = "cluster"
	ComputePath           = "compute/host"
	DatastorePath         = "datastore"
	DocumentPath          = "document"
	GroupPath             = "group"
	HookPath              = "hook"
	HostPath              = "host"
	ImagePath             = "image"
	InstanceGroupPath     = "instance-group"
	InstancePath          = "instance"
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

//go:generate go run "github.com/dmarkham/enumer" -type LCMState -linecomment -text

// LCMState is the Life Cycle Manager state of an instance.
type LCMState int

// LCMState values.
const (
	InitLCMState          LCMState = iota // init
	PrologLCMState                        // prolog
	BootLCMState                          // boot
	RunningLCMState                       // running
	MigrateLCMState                       // migrate
	SaveStopLCMState                      // save_stop
	SaveSuspendLCMState                   // save_suspend
	SaveMigrateLCMState                   // save_migrate
	PrologMigrateLCMState                 // prolog_migrate
	PrologResumeLCMState                  // prolog_resume
	EpilogStopLCMState                    // epilog_stop
	EpilogLCMState                        // epilog
	ShutdownLCMState                      // shutdown
	_
	_
	CleanupResubmitLCMState              // cleanup_resubmit
	UnknownLCMState                      // unknown
	HotplugLCMState                      // hotplug
	ShutdownPowerOffLCMState             // shutdown_poweroff
	BootUnknownLCMState                  // boot_unknown
	BootPowerOffLCMState                 // boot_poweroff
	BootSuspendedLCMState                // boot_suspended
	BootStoppedLCMState                  // boot_stopped
	CleanupDeleteLCMState                // cleanup_delete
	HotplugSnapshotLCMState              // hotplug_snapshot
	HotplugNICLCMState                   // hotplug_nic
	HotplugSaveAsLCMState                // hotplug_saveas
	HotplugSaveAsPowerOffLCMState        // hotplug_saveas_poweroff
	HotplutSaveAsSuspendedLCMState       // hotplug_saveas_suspended
	ShutdownUndeployLCMState             // shutdown_undeploy
	EpilogUndeployLCMState               // epilog_undeploy
	PrologUndeployLCMState               // prolog_undeploy
	BootUndeployLCMState                 // boot_undeploy
	HotplugPrologPowerOffLCMState        // hotplug_prolog_poweroff
	HotplugEpilogPowerOffLCMState        // hotplug_epilog_poweroff
	BootMigrateLCMState                  // boot_migrate
	BootFailureLCMState                  // boot_failure
	BootMigrateFailureLCMState           // boot_migrate_failure
	PrologMigrateFailureLCMState         // prolog_migrate_failure
	PrologFailureLCMState                // prolog_failure
	EpilogFailureLCMState                // epilog_failure
	EpilogStopFailureLCMState            // epilog_stop_failure
	EpilogUndeployFailureLCMState        // epilog_undeploy_failure
	PrologMigratePowerOffLCMState        // prolog_migrate_poweroff
	PrologMigratePowerOffFailureLCMState // prolog_migrate_poweroff_failure
	PrologMigrageSuspendLCMState         // prolog_migrate_suspend
	PrologMigrageSuspendFailureLCMState  // prolog_migrate_suspend_failure
	BootUndeployFailureLCMState          // boot_undeploy_failure
	BootStoppedFailureLCMState           // boot_stopped_failure
	PrologResumeFailureLCMState          // prolog_resume_failure
	PrologUndeployFailureLCMState        // prolog_undeploy_failure
	DiskSnapshotPowerOffLCMState         // disk_snapshot_poweroff
	DiskSnapshotRevertPowerOffLCMState   // disk_snapshot_revert_poweroff
	DiskSnapshotDeletePowerOffLCMState   // disk_snapshot_delete_poweroff
	DiskSnapshotSuspendLCMState          // disk_snapshot_suspend
	DiskSnapshotRevertSuspendedLCMState  // disk_snapshot_revert_suspended
	DiskSnapshotDeleteSuspendedLCMState  // disk_snapshot_delete_suspended
	DiskSnapshotLCMState                 // disk_snapshot
	_
	DiskSnapshotDeleteLCMState          // disk_snapshot_delete
	PrologMigrateUnknownLCMState        // prolog_migrate_unknown
	PrologMigrateUnknownFailureLCMState // prolog_migrate_unknown_failure
)

// Lock is the API payload based on the legacy xmlrpc backend.
type Lock struct {
	Locked bool      `json:"locked" yaml:"locked"`
	Owner  int       `json:"owner" yaml:"owner"`
	Time   time.Time `json:"time,omitempty" yaml:"time,omitempty"`
	ReqID  int       `json:"req_id" yaml:"req_id"`
}

// Permissions is the API payload based on the legacy xmlrpc backend.
type Permissions struct {
	Owner perms `json:"owner,omitempty" yaml:"owner,omitempty"`
	Group perms `json:"group,omitempty" yaml:"group,omitempty"`
	Other perms `json:"other,omitempty" yaml:"other,omitempty"`
}

type perms struct {
	Use    bool `json:"use,omitempty" yaml:"use,omitempty"`
	Manage bool `json:"manage,omitempty" yaml:"manage,omitempty"`
	Admin  bool `json:"admin,omitempty" yaml:"admin,omitempty"`
}

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
